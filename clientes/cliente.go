package main
import(
        "os"
        "strings"
        "io"
        "encoding/csv"
        "log"
        "fmt"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "github.com/432i/T1SisDistribuidos/logistica/chat"
)


type Retail struct {
        tipo string 
        id string 
        producto string
        valor string
        tienda string 
        destino string 
}
type Pyme struct{
        tipo string 
        id string
        producto string 
        valor string
        tienda string 
        destino string 
        prioritario string 
}
func cargarPyme() []Pyme{
        csvpyme, _ := os.Open("pymes.csv")
        readerpyme := csv.NewReader(csvpyme)
        var pedidospyme []Pyme
        for {
                lineapyme, error := readerpyme.Read()
                if error == io.EOF {
                        break
                }else if error != nil{
                        log.Fatal(error)
                }

                pedidospyme = append(pedidospyme, Pyme{
                        tipo: "pyme",
                        id: lineapyme[0],
                        producto: lineapyme[1],
                        valor: lineapyme[2],
                        tienda: lineapyme[3],
                        destino: lineapyme[4],
                        prioritario: lineapyme[5],
                })
        }
        fmt.Println(pedidospyme)
        return pedidospyme
}

func cargarRetail() []Retail{
        csvretail, _ := os.Open("retail.csv")
        readerretail := csv.NewReader(csvretail)
        var pedidosretail []Retail
        for {
                linearetail, error := readerretail.Read()
                if error == io.EOF {
                        break
                }else if error != nil{
                        log.Fatal(error)
                }
                pedidosretail = append(pedidosretail, Retail{
                        tipo: "retail",
                        id: linearetail[0],
                        producto: linearetail[1],
                        valor: linearetail[2],
                        tienda: linearetail[3],
                        destino: linearetail[4],
                })
        }
        fmt.Println(pedidosretail)
        return pedidosretail
}


func main(){
        var conn *grpc.ClientConn
        conn, err := grpc.Dial("10.6.40.149:9000", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()

        c := chat.NewChatServiceClient(conn)

        //response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
        //if err != nil {
        //        log.Fatalf("Error when calling SayHello: %s", err)
        //}
        //log.Printf("Response from server: %s", response.Body)
        
        fmt.Println("...\n")
        pedidosPyme := cargarPyme()
        pedidosRetail := cargarRetail()
        cantPyme := len(pedidosPyme)
        cantRetail := len(pedidosRetail)
        contPyme := 0
        contRetail := 0
        fmt.Println("csvs cargados correctamente\n")
        for{
                var respuesta string
                fmt.Println("ヾ(•ω•`)o Bienvenido ヾ(•ω•`)o, \n")
                fmt.Println("Ingrese la alternativa que desee: \n")
                fmt.Println("1 Enviar una orden desde una Pyme \n")
                fmt.Println("2 Enviar una orden desde el Retail \n")
                fmt.Println("3 Realizar seguimiento de un pedido \n")
                fmt.Println("432 para salir")
                _, err := fmt.Scanln(&respuesta)
                if err != nil {
                        fmt.Fprintln(os.Stderr, err)
                        return
                }
                fmt.Println("Tu respuesta fue: ")
                fmt.Println(respuesta)

                if strings.Compare("1", respuesta) == 0{
                        fmt.Println("XD1")
                        if contPyme == (cantPyme-1){
                                fmt.Println("No quedan más ordenes que realizar. Saliendo... \n ")
                        }else{
                                ordenPyme := pedidosPyme[contPyme]
                                message := chat.Orden{
                                        Tipo: ordenPyme.tipo,
                                        Id: ordenPyme.id,
                                        Producto: ordenPyme.producto,
                                        Valor: ordenPyme.valor,
                                        Tienda: ordenPyme.tienda,
                                        Destino: ordenPyme.destino,
                                        Prioritario: ordenPyme.prioritario,
                                }
                                response, err := c.EnviarOrden(context.Background(), &message)
                                log.Printf("Su codigo de seguimiento es %s", response.Body)
                                
                                contPyme := contPyme+1
                        }
                }
                if strings.Compare("2", respuesta) == 0{
                        fmt.Println("XD2")
                        if contRetail == (cantRetail-1){
                                fmt.Println("No quedan más ordenes que realizar. Saliendo... \n ")
                        }else{
                                ordenRetail := pedidosRetail[contRetail]
                                message := chat.Orden{
                                        Tipo: ordenRetail.tipo,
                                        Id: ordenRetail.id,
                                        Producto: ordenRetail.producto,
                                        Valor: ordenRetail.valor,
                                        Tienda: ordenRetail.tienda,
                                        Destino: ordenRetail.destino,
                                        Prioritario: "2",
                                }
                                response, err := c.EnviarOrden(context.Background(), &message)
                                log.Printf("Su codigo de seguimiento es %s", response.Body)

                                contRetail := contRetail+1
                        }

                }
                if strings.Compare("3", respuesta) == 0{
                        fmt.Println("X3D")
                }
                if strings.Compare("432", respuesta) == 0{
                        fmt.Println("X432D")
                        break
                }
        }
}