package handler

import "time"

func ordenarStructFriendsPosts(friendsPosts *[]struct{ Post, Date, Author string }) (err error) {

	for indx := 0; indx < len(*friendsPosts)-1; indx++ {

		for i := 0; i < len(*friendsPosts)-1; i++ {

			var t1, t2 time.Time

			t1, err = time.Parse(time.Stamp, (*friendsPosts)[i].Date)
			if err != nil {
				return
			}

			t2, err = time.Parse(time.Stamp, (*friendsPosts)[i+1].Date)
			if err != nil {
				return
			}

			if t1.Before(t2) {

				aux := (*friendsPosts)[i]
				(*friendsPosts)[i] = (*friendsPosts)[i+1]
				(*friendsPosts)[i+1] = aux

			}

		}

	}
	return
}
