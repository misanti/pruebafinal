package buyer

import "fmt"

func Delete(id int) error {
	for i, u := range buyers {
		if u.ID == id {
			buyers = append(buyers[:i], buyers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Buyer with ID '%v' not found", id)
}