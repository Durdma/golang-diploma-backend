package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type University struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DomainId  primitive.ObjectID `json:"domain_id" bson:"domain_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	ShortName string             `json:"short_name" bson:"short_name"`
	History   History            `json:"history" bson:"history"`
	Settings  Settings           `json:"settings" bson:"settings"`
}

type History struct {
	Body      string             `json:"body" bson:"body"`
	CreatedBy primitive.ObjectID `json:"created_by" bson:"created_by"`
	UpdatedBy primitive.ObjectID `json:"updated_by" bson:"updated_by"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Settings struct {
	ID                       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MainColor                string             `json:"main_color" bson:"main_color"`
	MainColorHover           string             `json:"main_color_hover" bson:"main_color_hover"`
	MainFooterFontColor      string             `json:"main_footer_font_color" bson:"main_footer_font_color"`
	MainFooterFontColorHover string             `json:"main_footer_font_color_hover" bson:"main_footer_font_color_hover"`
	MainFooterBgColor        string             `json:"main_footer_bg_color" bson:"main_footer_bg_color"`
	HeaderImage              string             `json:"header_image,omitempty" bson:"header_image,omitempty"`
	Label                    string             `json:"label" bson:"label"`
}

// Editor - Сущность редактора контента на сайте университета.
type Editor struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"` // id записи в MongoDB.
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Verification Verification       `json:"verification" bson:"verification"` // Статус верификации
	RegisteredAt time.Time          `json:"registered_at" bson:"registered_at"`
	LastVisitAt  time.Time          `json:"last_visit_at" bson:"last_visit_at"`
	UniversityID primitive.ObjectID `json:"university_id" bson:"university_id"`
	Session      Session            `json:"session" bson:"session"`
}
