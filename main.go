package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
)

type Aircraft struct{
	Name        string `json:"name"`
	Country     string `json:"country"`
	Speed       string `json:"speed"`
	Role        string `json:"role"`
	FirstFlight int    `json:"first_flight"`
	Type string `json:"type"`
}

var airy=[]Aircraft{
	{
    Name:        "B-2 Spirit",
    Country:     "USA",
    Speed:       "Mach 0.95",
    Role:        "Stealth Bomber",
    FirstFlight: 1989,
    Type:        "Bomber",
  },
  {
    Name:        "F-22 Raptor",
    Country:     "USA",
    Speed:       "Mach 2.25",
    Role:        "Air Superiority Fighter",
    FirstFlight: 1997,
    Type:        "Fighter",
  },
  {
    Name:        "Su-57",
    Country:     "Russia",
    Speed:       "Mach 2.0",
    Role:        "Multirole Stealth Fighter",
    FirstFlight: 2010,
    Type:        "Fighter",
  },
  {
    Name:        "Dassault Rafale",
    Country:     "France",
    Speed:       "Mach 1.8",
    Role:        "Multirole Fighter",
    FirstFlight: 1986,
    Type:        "Fighter",
  },
  {
    Name:        "J-20",
    Country:     "China",
    Speed:       "Mach 2.0",
    Role:        "Stealth Fighter",
    FirstFlight: 2011,
    Type:        "Fighter",
  },
  {
    Name:        "SR-71 Blackbird",
    Country:     "USA",
    Speed:       "Mach 3.3",
    Role:        "Strategic Reconnaissance",
    FirstFlight: 1964,
    Type:        "Reconnaissance",
  },
  {
    Name:        "Tu-160 White Swan",
    Country:     "Russia",
    Speed:       "Mach 2.05",
    Role:        "Strategic Heavy Bomber",
    FirstFlight: 1981,
    Type:        "Bomber",
  },
  {
    Name:        "Concorde",
    Country:     "UK / France",
    Speed:       "Mach 2.04",
    Role:        "Supersonic Passenger Airliner",
    FirstFlight: 1969,
    Type:        "Passenger",
  },
  {
    Name:        "Boeing 747-8",
    Country:     "USA",
    Speed:       "Mach 0.86",
    Role:        "Wide-body Passenger Airliner",
    FirstFlight: 2010,
    Type:        "Passenger",
  },
  {
    Name:        "Airbus A380",
    Country:     "Europe",
    Speed:       "Mach 0.85",
    Role:        "Double-deck Passenger Airliner",
    FirstFlight: 2005,
    Type:        "Passenger",
  },
  {
    Name:        "C-17 Globemaster III",
    Country:     "USA",
    Speed:       "Mach 0.74",
    Role:        "Strategic Tactical Cargo",
    FirstFlight: 1991,
    Type:        "Cargo",
  },
  {
    Name:        "An-225 Mriya",
    Country:     "Ukraine",
    Speed:       "Mach 0.70",
    Role:        "Strategic Ultra-heavy Cargo",
    FirstFlight: 1988,
    Type:        "Cargo",
  },
  {
    Name:        "Eurofighter Typhoon",
    Country:     "UK / Germany / Italy / Spain",
    Speed:       "Mach 2.0",
    Role:        "Multirole Fighter",
    FirstFlight: 1994,
    Type:        "Fighter",
  },
  {
    Name:        "C-5M Galaxy",
    Country:     "USA",
    Speed:       "Mach 0.79",
    Role:        "Strategic Heavy Cargo Airlifter",
    FirstFlight: 1968,
    Type:        "Cargo",
  },
  {
    Name:        "B-52H Stratofortress",
    Country:     "USA",
    Speed:       "Mach 0.84",
    Role:        "Strategic Long-range Bomber",
    FirstFlight: 1952,
    Type:        "Bomber",
  },
  {
    Name:        "Avro Vulcan",
    Country:     "UK",
    Speed:       "Mach 0.96",
    Role:        "Strategic Delta-wing Bomber",
    FirstFlight: 1952,
    Type:        "Bomber",
  },
  {
    Name:        "Saab JAS 39 Gripen",
    Country:     "Sweden",
    Speed:       "Mach 2.0",
    Role:        "Light Multirole Fighter",
    FirstFlight: 1988,
    Type:        "Fighter",
  },
  {
    Name:        "MiG-25 Foxbat",
    Country:     "Soviet Union",
    Speed:       "Mach 2.83",
    Role:        "Supersonic Interceptor",
    FirstFlight: 1964,
    Type:        "Fighter",
  },
  {
    Name:        "F-35A Lightning II",
    Country:     "USA",
    Speed:       "Mach 1.6",
    Role:        "Multirole Stealth Fighter",
    FirstFlight: 2006,
    Type:        "Fighter",
  },
  {
    Name:        "HAL Tejas",
    Country:     "India",
    Speed:       "Mach 1.6",
    Role:        "Delta-wing Light Combat Aircraft",
    FirstFlight: 2001,
    Type:        "Fighter",
  },
  {
    Name:        "An-124 Ruslan",
    Country:     "Soviet Union / Ukraine",
    Speed:       "Mach 0.70",
    Role:        "Heavy Strategic Cargo",
    FirstFlight: 1982,
    Type:        "Cargo",
  },
  {
    Name:        "Il-76 Candid",
    Country:     "Soviet Union / Russia",
    Speed:       "Mach 0.75",
    Role:        "Multi-purpose Strategic Cargo",
    FirstFlight: 1971,
    Type:        "Cargo",
  },
  {
    Name:        "B-1B Lancer",
    Country:     "USA",
    Speed:       "Mach 1.25",
    Role:        "Supersonic Strategic Bomber",
    FirstFlight: 1974,
    Type:        "Bomber",
  },
  {
    Name:        "Xian H-6N",
    Country:     "China",
    Speed:       "Mach 0.75",
    Role:        "Long-range Strategic Bomber",
    FirstFlight: 1959,
    Type:        "Bomber",
  },
  {
    Name:        "Airbus A350-1000",
    Country:     "Europe",
    Speed:       "Mach 0.85",
    Role:        "Long-range Passenger Airliner",
    FirstFlight: 2013,
    Type:        "Passenger",
  },
  {
    Name:        "Boeing 777-300ER",
    Country:     "USA",
    Speed:       "Mach 0.84",
    Role:        "Wide-body Passenger Airliner",
    FirstFlight: 1994,
    Type:        "Passenger",
  },
  {
    Name:        "Lockheed U-2",
    Country:     "USA",
    Speed:       "Mach 0.67",
    Role:        "High-altitude Reconnaissance",
    FirstFlight: 1955,
    Type:        "Reconnaissance",
  },
}

func airyhandler(w http.ResponseWriter,r *http.Request){
	rand.Seed(time.Now().UnixNano())
	randomAircraft:=airy[rand.Intn(len(airy))]
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(randomAircraft)
}

func main(){
	http.HandleFunc("/",airyhandler)
	fmt.Println("Server running on port 6969")
	http.ListenAndServe(":6969",nil)
}