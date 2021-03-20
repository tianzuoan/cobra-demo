package sql2struct

import (
	"fmt"
	"github.com/tianzuoan/cobra-demmo/internal/word"
	"os"
	"strings"
	"text/template"
)

const structTpl = `package model

type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}}{{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }}{{ if gt $typeLen 0 }}{{.Name | ToCamelCase}} {{.Type}} {{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

//生成结构体
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn, filePath string) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderScoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	var fileName string
	if strings.HasSuffix(filePath, "/") {
		fileName = filePath + tableName + ".go"
	} else {
		fileName = filePath + "/" + tableName + ".go"
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("创建结构体文件%s失败，err:%v", fileName, err)
	}
	err = tpl.Execute(file, tplDB)
	err = tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
