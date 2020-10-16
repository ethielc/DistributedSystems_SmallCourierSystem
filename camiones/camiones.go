package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"bufio"
	"context"
	"math/rand"
	"github.com/432i/T1SisDistribuidos/logistica/chat"
	"google.golang.org/grpc"
)

type Camion struct {
	Tipo string
	Paquete1 chat.Paquete
	Paquete2 chat.Paquete
}

func getTime() string {
    t := time.Now()
    return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
}

func Intento(paquete chat.Paquete) {
	if paquete.Tipo == "retail" {
		if paquete.Intentos < 3 {
			if rand.float64() <= 0.8 {
				paquete.Estado = "Recibido"
			}
			else {
				paquete.Intentos += 1
			}
		} else {
			paquete.Estado = "No Recibido"
		}
	} else {
		if paquete.Intentos * 10 < paquete.Valor && paquete.Intentos < 2 {
			if rand.float64() <= 0.8 {
				paquete.Estado = "Recibido"
			} else {
				paquete.Intentos += 1
			}
		} else {
			paquete.Estado = "No Recibido"
		}
	}
}

func Entrega(camion Camion) {
	if camion.Paquete1.Valor > camion.Paquete2.Valor {
		Intento(camion.Paquete1)
	} else {
		Intento(camion.Paquete1)
	}
}

func Carga(camion Camion) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.149:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	if err != nil {
		log.Fatalf("No se pudo generar comunicacion: %s", err)
	}

	mensaje := chat.Message{
		Body: camion.Tipo,
	}

	paquete1, _ := c.PaqueteQueueToCamion(context.Background(), &mensaje)
	camion.Paquete1 = chat.Paquete{
		Id:       paquete1.GetId(),
		Track:    paquete1.GetTrack(),
		Tipo:     paquete1.GetTipo(),
		Intentos: paquete1.GetIntentos(),
		Estado:   paquete1.GetEstado(),
	}

	paquete2, _ := c.PaqueteQueueToCamion(context.Background(), &mensaje)
	camion.Paquete2 = chat.Paquete{
		Id:       paquete2.GetId(),
		Track:    paquete2.GetTrack(),
		Tipo:     paquete2.GetTipo(),
		Intentos: paquete2.GetIntentos(),
		Estado:   paquete2.GetEstado(),
	}

	Entrega(camion)


}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el tiempo de espera de los camiones\n")
	tEspera, _ := strconv.Atoi(reader.ReadString('\n'))

	fmt.Println("Ingrese el tiempo de envio de los paquetes\n")
	tEnvio, _ := strconv.Atoi(reader.ReadString('\n'))

    CamionR1 := Camion {
		Tipo: "retail",
	}
	CamionR2 := Camion {
		Tipo: "retail",
	}
	CamionN := Camion{
		Tipo: "normal",
	}

	for {
		time.Sleep(tEspera * time.Second)
		go Carga(CamionR1)
		go Carga(CamionR2)
		go Carga(CamionN)

	}

}