package esepunittests

import (
	"testing"
)

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	// if expected_value != actual_value {
	// 	t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	// }
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	// if expected_value != actual_value {
	// 	t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	// }
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	// if expected_value != actual_value {
	// 	t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	// }
}

func TestGetGradeNegative(t *testing.T) {
	expected_value := "F" // need to put the letter grade

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", -2, Assignment)
	gradeCalculator.AddGrade("exam 1", -5, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", -8, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	// if expected_value != actual_value { // compares to a string
	// 	t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	// }
} // edge case for negative values since main code doesn't exclude out of bounds values

func TestGetGradeAPlus(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 120, Assignment)
	gradeCalculator.AddGrade("exam 1", 105, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 115, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	// if expected_value != actual_value {
	// 	t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	// }
} // edge case for above 100 (A)
