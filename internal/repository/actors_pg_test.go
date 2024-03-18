package repository

import (
	"film-library-service/configs"
	"film-library-service/internal/entity"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

func TestActorPG_GetAll(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	db, _ := NewPostgresDB(configs.NewConfig())
	tests := []struct {
		name    string
		fields  fields
		want    []entity.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{db: db},
			want:    []entity.Actors{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ActorPG{
				db: tt.fields.db,
			}
			_, err := r.GetAll()
			if err != nil {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestActorPG_GetMoviesByActorID(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		actorID int
	}
	db, _ := NewPostgresDB(configs.NewConfig())
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Movies
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{db: db},
			args:    args{actorID: 2},
			want:    []entity.Movies{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ActorPG{
				db: tt.fields.db,
			}
			_, err := r.GetMoviesByActorID(tt.args.actorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMoviesByActorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestActorPG_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		actor entity.Actors
	}
	db, _ := NewPostgresDB(configs.NewConfig())
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{db: db},
			args: args{
				actor: entity.Actors{
					LastName:   "test",
					FirstName:  "test",
					MiddleName: "test",
					GenderID:   2,
					DateBirth:  time.Now(),
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ActorPG{
				db: tt.fields.db,
			}
			_, err := r.Create(tt.args.actor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestActorPG_AddMovie(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		actorID int
		movieID int
	}
	db, _ := NewPostgresDB(configs.NewConfig())
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{db: db},
			args: args{
				actorID: 1,
				movieID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ActorPG{
				db: tt.fields.db,
			}
			if err := r.AddMovie(tt.args.actorID, tt.args.movieID); (err != nil) != tt.wantErr {
				t.Errorf("AddMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
