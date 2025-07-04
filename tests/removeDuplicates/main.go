package main

type Incident struct {
	Id       int64
	Incident string
}

func removeDuplicates(incidents []Incident) {
	seen := make(map[string]bool)
	toDelete := []int64{}

	for _, incident := range incidents {
		if !seen[incident.Incident] {
			seen[incident.Incident] = true
		} else {
			toDelete = append(toDelete, incident.Id)
		}
	}

	for _, id := range toDelete {
		println("видалення дублюючого інциденту з ID:", id)
	}
}

func main() {
	incidents := []Incident{
		{Id: 1, Incident: "A"},
		{Id: 2, Incident: "B"},
		{Id: 3, Incident: "A"},
		{Id: 4, Incident: "A"},
		{Id: 5, Incident: "C"},
		{Id: 6, Incident: "B"},
	}

	removeDuplicates(incidents)
}