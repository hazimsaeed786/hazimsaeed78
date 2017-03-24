## Сканирование документов

```golang
package main

import (

//        "fmt"
        "log"
         r "github.com/dancannon/gorethink"

)

// Declaration inetrfaces
type Mst map[string]interface{} // map - structure - interface
type Mif []interface{}

//
// Вставка из одной таблицы в другую
// Задание для себя
//
func main() {


   // Подключение к базе данных
   // session, err := r.Connect(r.ConnectOpts{Address: "10.0.20.5:28015", Database: "test"})
    session, err := r.Connect(r.ConnectOpts{Address: "195.128.18.66:28015", Database: "test"})

    // Обработка ошибок
    if err != nil {
        log.Fatalln(err)
       }
/*
       var objList = []interface{}{
     map[string]interface{}{"id": 1, "g1": 1, "g2": 1, "num": 0},
     map[string]interface{}{"id": 2, "g1": 2, "g2": 2, "num": 5},
     map[string]interface{}{"id": 3, "g1": 3, "g2": 2, "num": 10},
     map[string]interface{}{"id": 4, "g1": 2, "g2": 3, "num": 0},
     map[string]interface{}{"id": 5, "g1": 2, "g2": 3, "num": 100},
     map[string]interface{}{"id": 6, "g1": 1, "g2": 1, "num": 15},
     map[string]interface{}{"id": 7, "g1": 1, "g2": 2, "num": 0},
     map[string]interface{}{"id": 8, "g1": 4, "g2": 2, "num": 50},
     map[string]interface{}{"id": 9, "g1": 2, "g2": 3, "num": 25},
}

*/
//tt,_:=r.Db("test").Table("ts")

//var objList1 = []interface{}{map[string]interface{}{tt}}

   //r.Db('test').Table('tg').Insert(r.Db('test').Table('Calc').orderBy(r.desc('ids')).limit(100)).RunWrite(session)
   // r.Db("test").Table("tg").Delete().Run(session)

    //r.Db("test").Table("tg").Insert([]interface{}{Mst{r.Db("test").Table("ts").Without("id")}}).Run(session)
    //r.Db("test").Table("tg").Insert(Mst{"idss":"222"}).RunWrite(session)

    //r.Db("test").Table("tg").Insert(Mst{r.Db("test").Table("ts").Without("id")}).Run(session)

// tabb:= r.Db("test").Table("ts")
// Tabl:=[]interface{}{map[string]interface{}{tabb}}

    //var res []Mst{}
//var noDupNumObjList

  /*
type Stt struct{
         age int
         index int
    }
*/

//var ss = []Mst{}

//var res = map[string]interface{}{}
// Формирование для одного документа
    /*
     m := make(map[string]interface{})

     resp, errj := r.Db("test").Table("ts").Run(session)
    errj=resp.All(&m)

     // Обработка ошибок
     if errj != nil {
          panic("\n \n Problem with Load Document........................................!!!! \n \n")
     }

    */

/*
     var ObjLoadData = Mif{
          Mst{"id": 1, "Title": "Перемещение товара", "Description": "Перемещение медикаментов", "Gr": 0},
          Mst{"id": 2, "Title": "Расчет потребностей", "Description": "Расчет потребностей", "Gr": 0},
          Mst{"id": 3, "Title": "Данные для BI", "Description": "Выходные данные по BI", "Gr": 0},
          Mst{"id": 4, "Title": "Формирование заказа", "Description": "Автоматическое формирование заказа", "Gr": 0},
     }
*/

    //r.Db("test").Table("tt").Delete().Exec(session)
     rr,_:= r.Db("test").Table("tt").Run(session)
     zz := Mif{Mst{&rr}}

   r.Db("test").Table("tg").Insert(&zz).RunWrite(session)

//fmt.Println("Ok")
}
```
