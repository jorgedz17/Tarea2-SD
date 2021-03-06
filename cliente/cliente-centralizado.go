package main

import (
	pb "Lab2/Tarea2-SD/pipeline"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

/*-----------------------------------------------------------------------------------------*/
func pedir_archivo(opcion string) (int, string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist157:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	response, err := c.SolicitarUbicaciones(context.Background(), &pb.ConsultaUbicacion{NombreArchivo: opcion})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	partes := response.Partes
	ubicacion := response.Ubicaciones
	return int(partes), ubicacion
}

func howManyChunks(FileName string) (uint64, uint64) {
	fileToBeChunked := FileName
	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := uint64(fileInfo.Size())
	fileChunk := 256000 //Bytes
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	return totalPartsNum, fileSize
}

func sendChunk(partToSend int, bookName string, maquina string) {
	chunkToSend := bookName + "_" + strconv.FormatUint(uint64(partToSend), 10)
	chunkBytes, err := ioutil.ReadFile(chunkToSend) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	var conn *grpc.ClientConn
	conn, err1 := grpc.Dial(maquina+":50054", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	response, err := c.ClientToDataNode(context.Background(), &pb.DataChuck{Valor: int32(partToSend), Chunck: chunkBytes, NombreArchivo: bookName})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	fmt.Println("# DataNode responde: Se a recibido chunk numero ", response.Valor)
	return
}

/*----------------------------------------------------------------------------------------------------------------------------------------*/
// Esta función separa el archivo en diferentes archivos de 250 KB cada uno
/*----------------------------------------------------------------------------------------------------------------------------------------*/
func gutTheFile(fileName string) uint64 {
	fileChunk := 256000 //Bytes
	totalPartsNum, fileSize := howManyChunks(fileName)
	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(float64(fileChunk), float64(int64(fileSize)-int64(i*uint64(fileChunk)))))
		partBuffer := make([]byte, partSize)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println()
			os.Exit(1)

		}
		file.Read(partBuffer)
		fileName := fileName + "_" + strconv.FormatUint(i, 10)
		_, err1 := os.Create(fileName)
		if err != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		fmt.Println("Split to : ", fileName)
	}
	return totalPartsNum
}

/*-----------------------------------------------------------------------------------------*/
func requestChunk(maquina string, fileChunk int, bookTag string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina+":50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	fmt.Println("waiting >>>")
	fmt.Println("*     Chunk Solicitado      *")
	fmt.Println(fileChunk)
	fmt.Println("*****************************")
	c := pb.NewGreeterClient(conn)
	response, err := c.SayHello(context.Background(), &pb.Book{Request: int32(fileChunk), BookName: bookTag})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("La parte solicitada es: %d", response.Valor)
	fileName := bookTag + "_" + strconv.FormatUint(uint64(fileChunk), 10)
	fmt.Println("se recibe: ", fileName)
	ioutil.WriteFile(fileName, response.Chuck, os.ModeAppend)
}

func avisar_termino(maquina string, bookTag string, partes int) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina+":50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	aux_maquina := 0
	if maquina == "dist158" {
		aux_maquina = 0
	} else if maquina == "dist159" {
		aux_maquina = 1
	} else {
		aux_maquina = 2
	}
	log.Printf("Enviando aviso")
	response, err := c.YadaYada(context.Background(), &pb.ClientCheck{Request: int32(aux_maquina), BookName: bookTag, Partes: int32(partes)})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	if response.Valor == 0 {
		log.Printf("La distribucion fue exitosa")
	} else {
		log.Printf("La distribucion no fue posible")
	}
}

/*---------------------------------------------------*/
func stitchTheFile(originalName string, totalPartsNum uint64) {
	writePosition := int64(0)
	newFileName := "NEW_" + originalName
	file, e := os.Create(newFileName)
	if e != nil {
		log.Fatal(e)
	}

	for j := uint64(0); j < totalPartsNum; j++ {
		currentChunkFileName := originalName + "_" + strconv.FormatUint(j, 10)

		newFileChunk, err := os.Open(currentChunkFileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer newFileChunk.Close()

		chunkInfo, err := newFileChunk.Stat()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		chunkSize := chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)

		fmt.Println("Appending at position : [", writePosition, "] bytes")
		writePosition = writePosition + chunkSize

		reader := bufio.NewReader(newFileChunk)
		_, err = reader.Read(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		n, err := file.Write(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file.Sync()
		chunkBufferBytes = nil

		fmt.Println("Written ", n, " bytes")

		fmt.Println("Recombining part [", j, "] into : ", newFileName)
	}
	file.Close()
}

func archivos_disponibles() []string {
	var conn *grpc.ClientConn
	conn, err1 := grpc.Dial("dist157:50055", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	response, err := c.FilesAvl(context.Background(), &pb.Resultado{Valor: int32(1)})
	if err != nil {
		log.Fatalf("Error when calling FilesAvl: %s", err)
	}
	StrFiles := response.NombreArchivo
	files := strings.Split(StrFiles, "-")
	return files
}

func verificar_archivo(nombre_archivo string, archivos_dis []string) int {
	for cont := 0; cont < len(archivos_dis); cont++ {
		if archivos_dis[cont] == nombre_archivo {
			return 0
		}
	}
	return 1
}

func check_maquinas() int {
	m := [3]string{"dist158", "dist159", "dist160"}
	blag := 0
	for i := 0; i < len(m); i++ {
		blag = 0
		var conn *grpc.ClientConn
		mach := m[i] + ":50054"
		conn, err := grpc.Dial(mach, grpc.WithInsecure())
		if err != nil {
			fmt.Println("Maquina no disponible, distribucion rechazada")
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)
		response, err := c.TesteoEstado(context.Background(), &pb.Bla{Valor: int32(1)})
		if err != nil {
			fmt.Println("Maquina no disponible, distribucion rechazada")
			blag = 1
		}
		fmt.Println("Maquina Respondio ", response)
		if blag == 0 {
			return i
		}
		defer conn.Close()
	}
	fmt.Println("Todas las maquinas disponibles, distribucion aceptada")
	return -1
}

func solicitar_archivo() {
	archivos_dis := archivos_disponibles()
	opcion := "bandera"
	check := 1
	for check != 0 {
		fmt.Println("### Los archivos disponibles son los siguientes:")
		for i := 0; i < len(archivos_dis); i++ {
			fmt.Println(archivos_dis[i])
		}
		fmt.Println("Ingrese el nombre exacto de alguno de los archivos disponibles")
		fmt.Println("# Ejemplo: test.pdf ")
		fmt.Scanf("%s", &opcion)
		check = verificar_archivo(opcion, archivos_dis)
		if check == 0 {
			partes, maquinas := pedir_archivo(opcion)
			aux_maquina := strings.Split(maquinas, "-")
			totalChunks := uint64(partes)
			aux := 0
			for j := uint64(0); j < totalChunks; j++ {
				aux = int(j)
				requestChunk(aux_maquina[aux], aux, opcion)
			}
			stitchTheFile(opcion, totalChunks)
			fmt.Println("[°] Archivo Reconstruido y disponible")

		} else {
			fmt.Println("### Escriba un nombre de archivo valido")
		}
	}
}

func subir_archivo() {
	opcion := ""
	fmt.Println("# Ingrese el nombre exacto del archivo que va a subir")
	fmt.Println("# Ejemplo: test.pdf ")
	fmt.Scanf("%s", &opcion)
	fmt.Println("[°] Comenzando proceso para subir el archivo")
	cantidad_partes := gutTheFile(opcion)
	maquina_dis := check_maquinas()
	m := [3]string{"dist158", "dist159", "dist160"}
	maquina := m[maquina_dis]
	for i := 0; i < int(cantidad_partes); i++ {
		sendChunk(i, opcion, maquina)
	}
	fmt.Println("[°] Todos los chunks enviados")
	avisar_termino(maquina, opcion, int(cantidad_partes))
}

func menu() {
	opcion := 0
	fmt.Println("### Bienvenid@ a la tarea 2 de Distribuidos")
	for opcion != 3 {
		fmt.Println("Escriba 1 para poder subir un pdf")
		fmt.Println("Escriba 2 para poder pedir un pdf")
		fmt.Println("Escriba 3 para salir")
		fmt.Scanf("%d", &opcion)
		if opcion == 1 {
			subir_archivo()
		}
		if opcion == 2 {
			solicitar_archivo()
		}
	}

}

func main() {
	menu()
}
