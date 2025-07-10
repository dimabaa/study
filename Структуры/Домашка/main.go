package main

import "fmt"

type Job_seeker struct {
	// длина имени от 3 до 20 символов
	name string
	// возраст от 0 до 99 лет
	age int
	// male \ female
	gender string
	// длина поля до 100 символов
	profession string
	// student \ junior \ middle \ middle+ \ senior \ teamlead
	grade string
	// true \ false
	looking_job bool
	// 0 - 10.0
	rating float64
	// none \ a \ b \ c \ d \ e
	driver_license string
	// school \ middle \ high
	education string
}

func (js *Job_seeker) newName(newName string) {
	if len(newName) > 3 && len(newName) < 99 {
		js.name = newName
		fmt.Println("\nИмя изменено на:", js.name)
	} else {
		fmt.Println("\n>>> WARN:\n	Имя не изменено, т.к. не прошло проверку на длину")
	}
}

func (js *Job_seeker) newAge(newAge int) {
	if newAge > 0 && newAge < 99 {
		js.age = newAge
		fmt.Println("\nВозраст изменен на:", js.age)
	} else {
		fmt.Println("\n>>> WARN:\n	Возраст не изменён, т.к. не прошёл проверку")
	}
}

func (js Job_seeker) newGender(newGender string) {
	if newGender == "male" || newGender == "female" {
		js.gender = newGender
		fmt.Println("\nПол изменён на:", js.gender)
	} else {
		fmt.Println("\n>>> WARN:\n	Пол не изменён, т.к. не прошёл проверку")
	}
}

func (js *Job_seeker) newProfession(newProfession string) {
	if len(newProfession) > 0 && len(newProfession) < 100 {
		js.profession = newProfession
		fmt.Println("\nПрофессия изменена на:", js.profession)
	} else {
		fmt.Println("\n>>> WARN:\n	Профессия не изменена, т.к. наименование слишком длинное")
	}
}

func (js *Job_seeker) newGrade(newGrade string) {
	if newGrade == "student" ||
		newGrade == "junior" ||
		newGrade == "middle" ||
		newGrade == "middle+" ||
		newGrade == "senior" ||
		newGrade == "teamlead" {
		js.grade = newGrade
		fmt.Println("\nГрейд изменён на:", js.grade)
	} else {
		fmt.Println("\n>>> WARN:\n	Грейд не изменён, т.к. не входит в перечень допустимых:\n	student | junior | middle | middle+ | senior | teamlead")
	}
}

func (js *Job_seeker) ratingUP(rating float64) {
	if js.rating+rating <= 10.0 {
		js.rating += rating
		fmt.Println("\nРейтинг изменён на:", js.rating)
	} else {
		fmt.Println("\n>>> WARN:\n	Рейтинг не может быть больше 10")
	}
}

func (js *Job_seeker) ratingDOWN(rating float64) {
	if js.rating-rating > 0.0 {
		js.rating -= rating
		fmt.Println("\nРейтинг изменён на:", js.rating)
	} else {
		fmt.Println("\n>>> WARN:\n	Рейтинг не может быть меньше 0")
	}
}

func (js *Job_seeker) newLookingJob(newLookingJob bool) {
	if fmt.Sprintf("%T", newLookingJob) == "bool" {
		js.looking_job = newLookingJob
		fmt.Println("\nСтатус поиска работы изменён на:", js.looking_job)
	} else {
		fmt.Println("\n>>> WARN:\n	Статус поиска работы не изменён, не верный тип")
	}
}

func (js *Job_seeker) newDriverLic(newDriverLic string) {
	if newDriverLic == "none" ||
		newDriverLic == "a" ||
		newDriverLic == "b" ||
		newDriverLic == "c" ||
		newDriverLic == "d" ||
		newDriverLic == "e" {
		js.driver_license = newDriverLic
		fmt.Println("\nКатегория водительского удостоверения изменена на:", js.driver_license)
	} else {
		fmt.Println("\n>>> WARN:\n	Категория водительского удостоверения не изменена, т.к. не входит в перечень допустимых:\n	none | a | b | c | d | e")
	}
}

func (js *Job_seeker) newEducation(newEducation string) {
	if newEducation == "school" ||
		newEducation == "middle" ||
		newEducation == "high" {
		if js.education == newEducation {
			fmt.Println("\n>>> INFO:\n	Изменений в образовании не произведено, т.к. пришло аналогичное значение")
		} else {
			js.education = newEducation
			fmt.Println("\nОбразование изменено на:", js.education)
		}
	} else {
		fmt.Println("\n>>> WARN:\n	Образование не изменено, т.к. не входит в перечень допустимых:\n	school | middle | high")
	}
}

// конструктор для структуры

func NewJob_seeker(
	name string,
	age int,
	gender string,
	profession string,
	grade string,
	looking_job bool,
	rating float64,
	driver_license string,
	education string,
) Job_seeker {
	fmt.Println("Валидирую имя")
	if len(name) < 3 || len(name) > 99 {
		fmt.Println("Имя не прошло валидацию")
		return Job_seeker{}
	}
	fmt.Println("Валидирую возраст")
	if age < 0 || age > 99 {
		fmt.Println("Возраст не прошёл валидацию")
		return Job_seeker{}
	}
	fmt.Println("Валидирую пол")
	if gender != "female" && gender != "male" {
		fmt.Println("Пол не прошёл валидацию")
		return Job_seeker{}
	}
	return Job_seeker{
		name:           name,
		age:            age,
		gender:         gender,
		profession:     profession,
		grade:          grade,
		looking_job:    looking_job,
		rating:         rating,
		driver_license: driver_license,
		education:      education,
	}
}

func main() {
	job_seeker := NewJob_seeker(
		"Dmitry",
		33,
		"male",
		"teamlead support engeneer",
		"middle",
		true,
		9.0,
		"b",
		"high",
	)

	// вывод структуры в читаемом виде
	fmt.Println(
		"\n*** Содержимое структуры job_seeker ***",
		"\nИмя:", job_seeker.name,
		"\nВозраст:", job_seeker.age,
		"\nПол:", job_seeker.gender,
		"\nПрофессия:", job_seeker.profession,
		"\nГрейд:", job_seeker.grade,
		"\nИщет работу:", job_seeker.looking_job,
		"\nРейтинг:", job_seeker.rating,
		"\nВодительское удостоверение, категория:", job_seeker.driver_license,
		"\nОбразование:", job_seeker.education,
	)

	// применяем методы к структуре
	job_seeker.newName("Julia")
	job_seeker.newAge(18)
	job_seeker.newGender("female")
	job_seeker.newProfession("Golang developer")
	job_seeker.newGrade("middle++")
	job_seeker.newLookingJob(false)
	job_seeker.ratingUP(0.5)
	job_seeker.newDriverLic("y")
	job_seeker.newEducation("high")
	fmt.Println("\nОбновленные данные структуры\n	", job_seeker)
}
