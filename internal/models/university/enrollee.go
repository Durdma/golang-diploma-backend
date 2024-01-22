package university

import "go.mongodb.org/mongo-driver/bson/primitive"

// Enrollee - Сущность, описывающая абитуриента.
type Enrollee struct {
	ID      primitive.ObjectID // id записи в MongoDB.
	Name    string             // Имя абитуриента.
	Surname string             // Фамилия абитуриента.
	Code    string             // Код, присвоенный абитуриенту.
	Points  int                // Количество баллов, полученных за ЕГЭ или вступительные испытания.
}

// EnrolleeNews - Сущность для объявлений для абитуриентов.
type EnrolleeNews struct {
}

// EnrolleeTable - Сущность для конкурсного списка абитуриентов.
type EnrolleeTable struct {
}

// AdmissionRules - Сущность для описания правил приема в университет.
type AdmissionRules struct {
}

// EnrolleeInfo - Сущность для описания дополнительной информации для поступления в университет.
type EnrolleeInfo struct {
}

// ExamsInfo - Сущность для описания правил проведения вступительных экзаменов.
type ExamsInfo struct {
}
