package domain

import (
	api "back/src/generated"
	"errors"
	"strconv"
	"strings"
)

type ExistHoarderRecord struct {
	HoarderID    int    `db:"hoarder_id"`
	StatusID     int    `db:"status_id"`
	BookID       int    `db:"book_id"`
	Title        string `db:"title"`
	BookUserID   int    `db:"book_user_id"`
	TagIDArray   string `db:"tag_id_array"`
	TagNameArray string `db:"tag_name_array"`
}

func ConvertStautstoId(s api.Status) (int, error) {
	switch s {
	case api.Todo:
		return 1, nil
	case api.Wip:
		return 2, nil
	case api.Done:
		return 3, nil
	default:
		return -1, errors.New("invalid status")
	}
}

func convertIdtoStatus(i int) (api.Status, error) {
	switch i {
	case 1:
		return api.Todo, nil
	case 2:
		return api.Wip, nil
	case 3:
		return api.Done, nil
	default:
		return "", errors.New("invalid status")
	}
}

func (r *ExistHoarderRecord) ToExistHoarderBook() (api.ExistHoarderBook, error) {
	s, err := convertIdtoStatus(r.StatusID)
	if err != nil {
		return api.ExistHoarderBook{}, err
	}

	var tags []api.Tag
	if r.TagIDArray != "" && r.TagNameArray != "" {
		ids := strings.Split(r.TagIDArray, ",")
		names := strings.Split(r.TagNameArray, ",")
		for i, id := range ids {
			num_id, err := strconv.Atoi(id)
			if err != nil {
				return api.ExistHoarderBook{}, ErrInternal
			}
			tags = append(tags, api.Tag{Id: &num_id, Name: &names[i]})
		}
	}

	return api.ExistHoarderBook{
		Book: &api.ExistBook{
			Title:  &r.Title,
			UserId: &r.BookUserID,
			BookId: &r.BookID,
		},
		HoarderId: &r.HoarderID,
		Status:    &s,
		Tags:      &tags,
	}, nil
}
