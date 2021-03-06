package bqschema

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"google.golang.org/api/bigquery/v2"
)

var _ = Describe("ToStructs", func() {
	Context("when converting result rows to array of structs", func() {
		It("will fill an array of structs of simple types whos names match", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "A",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "B",
							Type: "FLOAT",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "C",
							Type: "STRING",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "D",
							Type: "BOOLEAN",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "2.0",
							},
							&bigquery.TableCell{
								V: "some",
							},
							&bigquery.TableCell{
								V: "false",
							},
						},
					},
				},
			}

			type test1 struct {
				A int
				B float64
				C string
				D bool
			}

			expectedResult := []test1{
				test1{
					A: 1,
					B: 2.0,
					C: "some",
					D: false,
				},
			}

			var dst []test1

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

		It("will fill an array of structs of simple types whos names no matter the casing", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "lower",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "UPPER",
							Type: "FLOAT",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "Title",
							Type: "STRING",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "camelCase",
							Type: "BOOLEAN",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "2.0",
							},
							&bigquery.TableCell{
								V: "some",
							},
							&bigquery.TableCell{
								V: "false",
							},
						},
					},
				},
			}

			type test2 struct {
				Lower     int
				UPPER     float64
				Title     string
				CamelCase bool
			}

			expectedResult := []test2{
				test2{
					Lower:     1,
					UPPER:     2.0,
					Title:     "some",
					CamelCase: false,
				},
			}

			var dst []test2

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

		It("will fill an array of structs of simple types whos names match the JSON tag", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "alpha",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "beta",
							Type: "FLOAT",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "gamma",
							Type: "STRING",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "delta",
							Type: "BOOLEAN",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "2.0",
							},
							&bigquery.TableCell{
								V: "some",
							},
							&bigquery.TableCell{
								V: "false",
							},
						},
					},
				},
			}

			type test1 struct {
				A int     `json:"alpha"`
				B float64 `json:"beta"`
				C string  `json:"gamma"`
				D bool    `json:"delta"`
			}

			expectedResult := []test1{
				test1{
					A: 1,
					B: 2.0,
					C: "some",
					D: false,
				},
			}

			var dst []test1

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

		It("will fill an array of structs of non standard types", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I64",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I32",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I16",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I8",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "F64",
							Type: "FLOAT",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "F32",
							Type: "FLOAT",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "2.0",
							},
							&bigquery.TableCell{
								V: "2.0",
							},
						},
					},
				},
			}

			type test3 struct {
				I64 int64
				I32 int32
				I16 int16
				I8  int8
				F64 float64
				F32 float32
			}

			expectedResult := []test3{
				test3{
					I64: 1,
					I32: 1,
					I16: 1,
					I8:  1,
					F64: 2.0,
					F32: 2.0,
				},
			}

			var dst []test3

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

		It("will fill an array of structs of unsigned ints", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I64",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I32",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I16",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I8",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "required",
							Name: "I",
							Type: "INTEGER",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
							&bigquery.TableCell{
								V: "1",
							},
						},
					},
				},
			}

			type test4 struct {
				I64 uint64
				I32 uint32
				I16 uint16
				I8  uint8
				I   uint
			}

			expectedResult := []test4{
				test4{
					I64: 1,
					I32: 1,
					I16: 1,
					I8:  1,
					I:   1,
				},
			}

			var dst []test4

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

		It("will handle nullable fields", func() {
			response := &bigquery.QueryResponse{
				Schema: &bigquery.TableSchema{
					Fields: []*bigquery.TableFieldSchema{
						&bigquery.TableFieldSchema{
							Mode: "NULLABLE",
							Name: "i",
							Type: "INTEGER",
						},
						&bigquery.TableFieldSchema{
							Mode: "NULLABLE",
							Name: "s",
							Type: "STRING",
						},
						&bigquery.TableFieldSchema{
							Mode: "NULLABLE",
							Name: "f",
							Type: "FLOAT",
						},
						&bigquery.TableFieldSchema{
							Mode: "NULLABLE",
							Name: "b",
							Type: "BOOLEAN",
						},
					},
				},
				Rows: []*bigquery.TableRow{
					&bigquery.TableRow{
						F: []*bigquery.TableCell{
							&bigquery.TableCell{
								V: nil,
							},
							&bigquery.TableCell{
								V: nil,
							},
							&bigquery.TableCell{
								V: nil,
							},
							&bigquery.TableCell{
								V: nil,
							},
						},
					},
				},
			}

			type test5 struct {
				I int
				S string
				F float32
				B bool
			}

			expectedResult := []test5{
				test5{
					I: 0,
					S: "",
					F: 0.0,
					B: false,
				},
			}

			var dst []test5

			err := ToStructs(response, &dst)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(expectedResult, dst)).To(BeTrue())
		})

	})
})
