package main

import "testing"

func TestReplData(t *testing.T) {
	testData := "---\nmenu:\n    before:\n        name: tasks\n        weight: 5\ntitle: Обновление данных в реальном времени\n---\n\n# Задача: Обновление данных в реальном времени\n\nНапишите воркер, который будет обновлять данные в реальном времени, на текущей странице.\nТекст данной задачи менять нельзя, только время и счетчик.\n\nФайл данной страницы: `/app/static/tasks/_index.md`\n\nДолжен меняться счетчик и время:\n\nТекущее время: 2021-10-13 15:00:00\n\nСчетчик: 0\n\n\n\n## Критерии приемки:\n- [ ] Воркер должен обновлять данные каждые 5 секунд\n- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд\n- [ ] Время должно обновляться каждые 5 секунд"

	type args struct {
		oldData []byte
	}
	tests := []struct {
		name    string
		args    args
		Notwant string
	}{
		{"Test Replase Data", args{oldData: []byte(testData)}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplData(tt.args.oldData); got == tt.Notwant {
				t.Errorf("ReplData() = %v, want %v", got, tt.Notwant)
			}
		})
	}
}
