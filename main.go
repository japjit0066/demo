package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const conferenceTickets int = 50

// var conferenceName = "Go Conference"
// var remainingTickets uint = 50
// var bookings = make([]UserData, 0)

// type UserData struct {
// 	firstName       string
// 	lastName        string
// 	email           string
// 	numberOfTickets uint
// }

// var wg = sync.WaitGroup{}

// func main() {

// 	greetUsers()

// 	// for {
// 	firstName, lastName, email, userTickets := getUserInput()
// 	isValidName, isValidEmail := ValidateUserInput(firstName, lastName, email)

// 	if isValidName && isValidEmail {

// 		bookTicket(userTickets, firstName, lastName, email)

// 		wg.Add(1)
// 		go sendTicket(userTickets, firstName, lastName, email)

// 		firstNames := getFirstNames()
// 		fmt.Printf("The first names of bookings are: %v\n", firstNames)

// 		if remainingTickets == 0 {
// 			// end program
// 			fmt.Println("Our conference is booked out. Come back next year.")
// 			// break
// 		}
// 	} else {
// 		if !isValidName {
// 			fmt.Println("first name or last name you entered is too short")
// 		}
// 		if !isValidEmail {
// 			fmt.Println("email address you entered doesn't contain @ sign")
// 		}

// 	}
// 	//}
// 	wg.Wait()
// }

// func greetUsers() {
// 	fmt.Printf("Welcome to %v booking application\n", conferenceName)
// 	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		firstNames = append(firstNames, booking.firstName)
// 	}
// 	return firstNames
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets = remainingTickets - userTickets

// 	var userData = UserData{
// 		firstName:       firstName,
// 		lastName:        lastName,
// 		email:           email,
// 		numberOfTickets: userTickets,
// 	}

// 	bookings = append(bookings, userData)
// 	fmt.Printf("List of bookings is %v\n", bookings)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }

// func sendTicket(userTickets uint, firstName string, lastName string, email string) {
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
// 	time.Sleep(05 * time.Second)
// 	fmt.Println("#################")
// 	time.Sleep(03 * time.Second)
// 	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
// 	time.Sleep(02 * time.Second)
// 	fmt.Println("#################")

// 	fmt.Println("do you want to enter are newsletter (yes or no)")
// 	var yesNo = ""
// 	fmt.Scan(&yesNo)
// 	if yesNo == "no" {
// 		fmt.Println("ok bye see you next time")

// 	} else {
// 		var phoneNumber = ""

// 		fmt.Println("then give us your phone number so we will enter you in the newsletter")

// 		fmt.Scan(&phoneNumber)
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Printf("+ 39 %v \n", phoneNumber)

// 		fmt.Println("thank you so much for entering the newsletter ")
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println("byee! see you next time")
// 	}
// 	time.Sleep(600 * time.Millisecond)
// 	fmt.Println("⚠️ ⚠️ ⚠️", "", "system shut down", "⚠️ ⚠️ ⚠️")

// 	wg.Done()
// }

// var Name = "jscds corp"
// var Cds = 5

// func main() {
// 	var choosen_song string
// 	var songcd [6]string
// 	choosen_song = songcd[0]
// 	var email = ""

// 	var input = ""
// 	var yesNo = ""
// 	var Cdit = [5]string{"playa,", "napoletano,", "gelosa,", "piove,", "bebe"}
// 	switch hour := time.Now().Hour(); {
// 	case hour <= 12:
// 		fmt.Println("boungiorno")
// 	case hour <= 17:
// 		fmt.Println("boun pomeriggio")
// 	default:
// 		fmt.Println("bouna sera")
// 	}
// 	fmt.Printf("benvenuto a  %v\n", Name)
// 	time.Sleep(05 * time.Second)
// 	fmt.Println("scrivi la tua email :")
// 	fmt.Scan(&email)
// 	time.Sleep(05 * time.Second)
// 	fmt.Printf("noi abbiamo %v Cds in magazino\n", Cds)
// 	time.Sleep(02 * time.Second)
// 	fmt.Println("ecco qualli sono:")
// 	time.Sleep(05 * time.Second)
// 	fmt.Println(Cdit)
// 	time.Sleep(05 * time.Second)
// 	fmt.Println("segli una canzone che vuoli compare:")
// 	fmt.Scan(&choosen_song)
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Printf("ecco cosa hai ordinato : %v \n ", choosen_song)
// 	time.Sleep(5 * time.Millisecond)

// 	fmt.Printf("sei sicuro che voi comprare  %v (si o no) \n", choosen_song)
// 	fmt.Scan(&yesNo)
// 	if yesNo == "no" {
// 		fmt.Println("stiamo per cancelare il tuo ordine ")

// 		fmt.Println("ciao  !!")
// 		os.Exit(0)
// 	}
// 	fmt.Println("vuoi leggere il testo di", choosen_song, "(si or no)")
// 	fmt.Scan(&input)
// 	if input == "no" {
// 		fmt.Println("ok, andiamo avanti")
// 	} else {
// 		if choosen_song == "playa" {
// 			for i := 0; i < 30; i++ {
// 				time.Sleep(500 * time.Millisecond) // mimics work
// 				printProgressBar(i+1, 30, "caricamento", "", 25, "=")
// 			}
// 			fmt.Println("Agora o 212 VIP exala no corpo do pai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Exala no corpo do pai, do pai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Shawty, que pasa?")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ho aspettato troppo, ti ho aspettato fuori casa")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eri così bella, mi sembravi una galassia")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Una come te non si vede mai per la strada, yeah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Baby, non so dirti perché, tanto succederà di nuovo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tornerà tutto come prima e poi ti scorderai di me")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non fare così, sei la mia ghetto bitch")
// 			time.Sleep(500 * time.Millisecond)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ehi, bih, passo lì, sei la mia queen")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("(Down, down, down, down, down)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ehi, bih, non far così, sto coi miei G")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("(Ma ora sono in down, down, down, down, down)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Dimmi perché, baby, non stiamo insieme")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non so più come amarti ormai, no, mai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Baby, sai che voglio solo te")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Baby, sai che voglio solo te")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("So che vuoi scappare da guai, ma tanto siamo dentro ormai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sarà tardi se chiamerai, è inutile, non mi troverai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non va bene, non va bene mai")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ehi, bih, passo lì, sei la mia queen")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("(Down, down, down, down, down)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ehi, bih, non far così, sto coi miei G")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("(Ma ora sono in down, down, down, down, down)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma ora sono in down, down, down, down, down")
// 			fmt.Println("#############################################################################")

// 		} else if choosen_song == "napoletano" {
// 			for i := 0; i < 30; i++ {
// 				time.Sleep(500 * time.Millisecond) // mimics work
// 				printProgressBar(i+1, 30, "caricamento", "", 25, "=")
// 			}
// 			fmt.Println("Ero soltanto picciril quando un napoletano ma rat na pistol")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Spar ngap a tutt sti infam")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Diceva a tutt, Guard a chist, iss è di Milano")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ij nun sto buon cu la cap, ci vediamo e vediamo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ero soltanto picciril quando un napoletano ma rat na pistol")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Spar ngap a tutt sti infam")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Diceva a tutt, Guard a chist, iss è di Milano")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ij nun sto buon cu la cap, ci vediamo e vediamo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ci vediamo e vediamo, poi non ti rivediamo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ti mettiamo in cofano, poi ti portiamo in alfano")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Primi soldi in mano, Piazza Affari, Bergamo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tu impari, io imparo, a dodic'anni ho imparato nu cazz")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sto in mezzo ville, la rue la vrai, quel che dico lo fa Basta guardare la mia faccia, non quello che fa")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Da noi uno scemo quello che in zona vostra è uno (u scem)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vatten, pesc mocc, vatten")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Nik mok, vatten prima che vattim")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vatten, tu sei solo chiattill")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Culo grosso, chiatton (la tua tipa bocchin), ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La tua p niente fluss, niente p")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Il tuo rapper del cazzo niente fluss, no makatussi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tossici del cazzo, rapper del cazzo illusi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi state tutti sul cazzo, meritate solo abusi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Poi è inutile che vieni e ti scusi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Perché Dio perdona, ma la strada no")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Siamo stati ventiquattro mesi chiusi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tutti e ventiquattro le ho passate in mbouck")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Le ho passate con la guardia che mi portava là")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Che mi portava il")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fumo con forse una, me lo succhia una (ah, ah, ah, ah)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Avevo dodici quando il 112")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prese me e il mio amico per un furto di alcolici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Avevo tredici quando il 113")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prese solo il mio amico perché non potevan crederci")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non potevan crederci fin quando ho fatto sedici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non potevan credere che è un dodicenne a vendere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Facce di un bravo uaglione, questi sbirri pecore")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vi ho fatto già da bambino, ora sto a godere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ero soltanto picciril quando un napoletano ma rat na pistol")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Spar ngap a tutt sti infam")
// 			fmt.Println("Diceva a tutt, Guard a chist, iss è di Milano")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ij nun sto buon cu la cap, ci vediamo e vediamo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ero soltanto picciril quando un napoletano ma rat na pistol")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Spar ngap a tutt sti infam")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Diceva a tutt, Guard a chist, iss è di Milano")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ij nun sto buon cu la cap, ci vediamo e vediamo")
// 			fmt.Println("#############################################################################")
// 		} else if choosen_song == "gelosa" {
// 			for i := 0; i < 30; i++ {
// 				time.Sleep(500 * time.Millisecond) // mimics work
// 				printProgressBar(i+1, 30, "caricamento", "", 25, "=")
// 			}
// 			fmt.Println("Stasse non dirmi di no, na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Stai con me dopo lo show, na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Voglio che posi l'iPhone, na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Stasera possiamo bere e andare da me")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Dopo spegnere la luce e guardare la TV")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non fare quella gelosa, na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("So che vuoi dire qualcosa, lo stai facendo apposta")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vorrei poter ridarti tutte le lacrime")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Penso che il diavolo abbia messo gli occhi su di me")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non so come resistere al tuo sguardo in questi giorni")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non ho tempo per seguire tutti questi sogni")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La mia anima è a metà, na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La coperta è ancora vuota")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La riempio con il fumo e con le bevande viola")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mentre tu ti fai le storie")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ho paura che non senti i tuoi rimorsi nella notte")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vorrei incatenarti il cuore")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non lasciarlo più partire")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Dormo con la Glock sotto al cuscino")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Da quando non sei vicino")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ti compro qualsiasi cosa, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("No, non chiedo quanto costa, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Chiamate senza risposta, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sai che non lo faccio apposta, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E tu che fai la gelosa")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Dici che non me ne importa")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ho il tuo sapore in bocca")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non dirmi che è l'ultima volta")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Babe non sbirciare dentro i miei DM")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Puoi fare la gelosa, ma col tuo boyfriend")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sai che dove vado, schiaccio sì come in Space Jam")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E non metto niente meno che Isabel o Vivienne")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("È da tempo che mi chiedi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Una canzone che parla di te")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Me la scrivo stanotte dall'ultimo piano sorvolando tutta la cité")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Dopo solo tre bicchieri ti ho baciata poi si è fatta l'alba")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eravamo solo fatti, mica fatti l'uno per l'altra")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lo vedo da come parli che qualcosa non va")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vola via con me a Parigi o Londra")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Se vuoi il cielo, il cielo sia, stoppa questa gelosia")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na na na na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ti compro qualsiasi cosa, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("No, non chiedo quanto costa, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Chiamate senza risposta, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sai che non lo faccio apposta, na, na, na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E tu che fai la gelosa, dici che non me ne importa")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ho il tuo sapore in bocca")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non dirmi che è l'ultima volta")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eh eh eh eh")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na na na na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eh eh eh eh")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na na na na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eh eh eh eh")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na na na na")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Eh eh eh eh")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Na na na na")
// 			fmt.Println("#############################################################################")
// 		} else if choosen_song == "piove" {
// 			for i := 0; i < 30; i++ {
// 				time.Sleep(500 * time.Millisecond) // mimics work
// 				printProgressBar(i+1, 30, "caricamento", "", 25, "=")
// 			}
// 			fmt.Println("Diego")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("333 Mob")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay, piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay Zzala")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non faccio pose")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Io non ho iniziato nada ho solo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Finito cose")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Spero 'sto giro 'sti bimbi sperduti facciano i lupi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Guardi lo show, fra', che lasciamo buchi, lasciami il booking")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay, lascia giù l'ascia di guerra")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Faccia di merda, ti dico: Basta")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non regolare più l'asta del micro, puoi mettere il micro all'asta")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Musical mafia come Frank Sinatra")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lei fa il pieno, benzinaia")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fra' non è Balenci-, ma bensì Prada")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fra' tu parli e pensi nada")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lancio soldi come un freesbee")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E tu corri come un inning")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sei una thottie con il lipstick")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sono toxic come Britney")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Marinero, marinero")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Parli, parli, e vali zero")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sai che in giro c'ho più di un amico che si fa il dinero")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Con un uovo in tasca come Calimero")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay, piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole (Money Gang)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Piovono euro, damn")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lasciamo a casa l'ombrello")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La tua ragazza che mi sbava dietro")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La porto in bagno le sbavo il rossetto, damn, ehi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("S'è fatto grande il chico di Cinisello")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tu, mi spiace, non sali di livello (brr)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sali e scendi come droga all'eccesso (brr, brr, brr, brr)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lo voglio e lo prendo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Corro un quarto di miglio, c'ho un quarto di millie sul polso (skrrt, skrrt)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vengo da dove basta uno sguardo e sanno cos'hai addosso (Bu-bu-bu-bu)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non parlarmi ce l'ho scritto in faccia che ho l'umore storto")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Qualcuno mi vuole morto")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi sono tolto tutti gli invidiosi e gli infami di torno (pa-pa-pa-pa-pa)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tu non lo hai mai visto un millie, non hai mai visto un chilo (brr)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mai visto piangere mamma il giorno dell'affitto (brr, brr, brr)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ottocento cavalli non bastano se siamo in giro (skrrt)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuoi fare il capo ma sei un cavallino")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay, piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Okay piove, lo dice un brother")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vuol dire muoviti pure se fuori")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Prendono il sole")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fanculo dodici")
// 			fmt.Println("#############################################################################")
// 		} else if choosen_song == "bebe" {
// 			for i := 0; i < 30; i++ {
// 				time.Sleep(500 * time.Millisecond) // mimics work
// 				printProgressBar(i+1, 30, "caricamento", "", 25, "=")
// 			}
// 			fmt.Println("Oh, mah, mah, mah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Oh, mah, mah, mah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, dimmi che succede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vedrai tutto si risolve, basta che restiamo insieme")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non va bene  ")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tuo padre quanto beve, so che ha la mano pesante")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E che piangi tutte le sere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, voglio andare dal quartiere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma sembro condannato a morire fra le sirene")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fuck carabiniere, ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi dispiace, babe, ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("I miei occhi sono tristi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Il sole e la luna stasera fanno un eclissi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sento come piangi mentre mi vengon le crisi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi verso del whisky")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sono incasinato come un quadro di Kandinskij")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma so che mi capisci, infondo siamo uguali")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tu piangi e mi fissi mentre sto rischiando gli anni")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("È l'unico modo per andar via dai palazzi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ci metterò poco, non devo far passi falsi")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Sotto lo stesso cielo, da bimbo non piangevo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tutt'oggi sai che tengo tutte le cose dentro")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lontani da tuo padre, un posto segreto")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Te lo prometto, ci andremo presto")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non faccio raggaeton-ton")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Voglio sdrammatizzare per questo un po' ci provo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("La mia voce malinconica è un carillon-on")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Parlo con me stesso, non riesco a restare sobrio")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi sento da solo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, dimmi che succede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vedrai tutto si risolve, basta che restiamo insieme")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non va bene")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tuo padre quando bene so che ha la mano pesante")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E che piangi tutte le sere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, voglio andare dal quartiere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma sembro condannato a morire fra le sirene")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fuck carabiniere, ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi dispiace, babe, ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, ogni sera beve")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Suo padre in galera, sua madre manco la vede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Si trucca la sera, esce sola, marciapiede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Su un Panamera, a volta sopra un Mercedes")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ueh, ueh, ueh-ueh")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Nessuno la vede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Nessuno gli crede se racconta ciò che vede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Nessuno la vede, quindi nessuno ci crede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("È tutto un dare e avere, ma tutti vogliono avere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma non sanno che sei mia e io mai stato il tuo")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Lettere al Beccaria scritte solo in cella al buio")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ora che sono free-free penso ancora a te-te")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Viviamo come un film-film solo io e te-te (oh, mah, mah, mah)")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, dimmi che succede")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Vedrai tutto si risolve, basta che restiamo insieme")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Non va bene")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Tuo padre quando bene so che ha la mano pesante")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("E che piangi tutte le sere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Bebe, voglio andare dal quartiere")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Ma sembro condannato a morire fra le sirene")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Fuck carabiniere, ah")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Mi dispiace, babe, ah")
// 			fmt.Println("#############################################################################")
// 		} else {
// 			fmt.Println("questa canzone non esiste")
// 		}
// 	}
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("vi mandiamo una email ")

// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Printf("abbiamo mandato una email a %v \n", email)
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("grazie per aver comprato da", Name)
// 	fmt.Println("alla prossima volta ")
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("#######################")

// 	time.Sleep(200 * time.Millisecond)

// 	fmt.Println("⚠️ ⚠️ ⚠️", "", "system shut down", "⚠️ ⚠️ ⚠️")

// }
// func printProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {
// 	percent := float64(iteration) / float64(total)
// 	filledLength := int(length * iteration / total)
// 	end := ">"

// 	if iteration == total {
// 		end = "="
// 	}
// 	if iteration == total {
// 		suffix = "✅"
// 	}
// 	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
// 	fmt.Printf("\r%s [%s] %f%% %s", prefix, bar, percent, suffix)
// 	if iteration == total {
// 		fmt.Println()
// 	}
// }

var yesNo string

func main() {
	var email string
	var noNo string
	var companyName = "levissima"
	var water = ""
	var name string
	var number int
	var Input string
	fmt.Println("welcome to", companyName)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("enter your first name :")
	fmt.Scan(&name)
	time.Sleep(200 * time.Millisecond)
	for {
		fmt.Println("we have three type of water")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("flavoured  , sparkling  , natural ")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("choose your type of water ")
		fmt.Scan(&water)
		switch water {
		case "flavoured":

			rand.Seed(time.Now().UnixNano())

			fmt.Println()
			fmt.Println("this is flavoured water")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("do you want to buy flavoured water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok bye see you next time ")

			} else {

				fmt.Println("ok enter your email please : ")
				fmt.Scan(&email)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v flavoured water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(200 * time.Millisecond)
				fmt.Println("see you next time ", name)

			}
		case "sparkling":
			var yesNo string
			fmt.Println("this is sparkling water")
			fmt.Println("do you want to buy sparkling water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok bye see you next time ")

			} else {
				var email string
				fmt.Println("ok enter your email please : ")
				fmt.Scan(&email)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v sparkling water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(1200 * time.Millisecond)
				fmt.Println("see you next time ", name)

			}
		case "natural":
			var yesNo string
			fmt.Println("this is natural water")
			fmt.Println("do you want to buy natural water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok bye see you next time ")

			} else {
				var email string
				fmt.Println("ok enter your email please : ")
				fmt.Scan(&email)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v natural water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(1200 * time.Millisecond)

			}
		}

		fmt.Println("do you want to buy more water ? (yes or no)")
		fmt.Scan(&noNo)
		if noNo == "no" {
			fmt.Println("ok then bye see you next time", name)
			break
		} else {
			fmt.Println("do you want to use the same email that is ", email, "(yes or no)")
			fmt.Scan(&Input)
			if Input == "no" {

			} else {
				for {
					fmt.Println("we will use this email for your order", email)
					fmt.Println("we have three type of water")
					time.Sleep(200 * time.Millisecond)
					fmt.Println("flavoured  , sparkling  , natural ")
					time.Sleep(200 * time.Millisecond)
					fmt.Println("choose your type of water ")
					fmt.Scan(&water)
					switch water {
					case "flavoured":

						rand.Seed(time.Now().UnixNano())

						fmt.Println()
						fmt.Println("this is flavoured water")
						time.Sleep(300 * time.Millisecond)
						fmt.Println("do you want to buy flavoured water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok bye see you next time ")

						} else {
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v flavoured water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(200 * time.Millisecond)
							fmt.Println("see you next time ", name)

						}
					case "sparkling":
						var yesNo string
						fmt.Println("this is sparkling water")
						fmt.Println("do you want to buy sparkling water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok bye see you next time ")

						} else {
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v sparkling water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(1200 * time.Millisecond)
							fmt.Println("see you next time ", name)

						}
					case "natural":
						var yesNo string
						fmt.Println("this is natural water")
						fmt.Println("do you want to buy natural water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok bye see you next time ")

						} else {
							var email string
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v natural water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(1200 * time.Millisecond)
						}
					}

					fmt.Println("do you want to buy more water ? (yes or no)")
					fmt.Scan(&noNo)
					if noNo == "no" {
						fmt.Println("ok then bye see you next time", name)
						break
					} else {
						fmt.Println("do you want to use the same email that is ", email, "(yes or no)")
						fmt.Scan(&Input)
						if Input == "no" {
							break
						} else {
							continue
						}

					}
				}
			}
		}

	}
}

// func main() {
// 	var pl = fmt.Println
// 	i := 1
// 	for ; i <= 100; i++ {

// 		if (i >= 1) && (i <= 18) {
// 			pl(i, "", "important birthday")
// 		} else if (i == 21) || (i == 50) {
// 			pl(i, "", "important birthday ")
// 		} else if (i >= 80) && (i <= 100) {
// 			pl(i, "", "really important birthday")
// 		} else {
// 			pl(i, "", "not important birthday")
// 		}

// 	}
// }
