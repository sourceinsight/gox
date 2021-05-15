package time2

import (
	"reflect"
	"testing"
	"time"
)

func TestAddDay(t *testing.T) {
	type args struct {
		t   time.Time
		day int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test AddDay",
			args: args{
				t:   time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
				day: 1,
			},
			want: time.Date(2021, 10, 11, 10, 10, 10, 10, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDay(tt.args.t, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddHour(t *testing.T) {
	type args struct {
		t    time.Time
		hour int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test AddHour",
			args: args{
				t:    time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
				hour: 1,
			},
			want: time.Date(2021, 10, 10, 11, 10, 10, 10, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddHour(tt.args.t, tt.args.hour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddMinute(t *testing.T) {
	type args struct {
		t       time.Time
		minutes int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test AddMinute",
			args: args{
				t:       time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
				minutes: 1,
			},
			want: time.Date(2021, 10, 10, 10, 11, 10, 10, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMinute(tt.args.t, tt.args.minutes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddSecond(t *testing.T) {
	type args struct {
		t       time.Time
		seconds int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test AddSecond",
			args: args{
				t:       time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
				seconds: 1,
			},
			want: time.Date(2021, 10, 10, 10, 10, 11, 10, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSecond(tt.args.t, tt.args.seconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayEnd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test DayEnd",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 10, 10, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayEnd(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayStart(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test DayStart",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayStart(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHourEnd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test HourEnd",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 10, 10, 10, 59, 59, int(time.Second-time.Nanosecond), time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HourEnd(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HourEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHourStart(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test HourStart",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HourStart(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HourStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthEnd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test MonthEnd",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthEnd(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthStart(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test MonthStart",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthStart(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearEnd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test YearEnd",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 12, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearEnd(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YearEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearStart(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test YearStart",
			args: args{
				t: time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC),
			},
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearStart(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YearStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
