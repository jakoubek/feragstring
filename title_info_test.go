package feragstring

import "testing"

func TestTitleInfo_PrintObjectName(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetPrintObjectName("FERAGTITLE")

	want := "+93FERAGTITLE  "

	if ti.PrintObjectName() != want {
		t.Errorf("PrintObjectName = %s, wanted %s", ti.PrintObjectName(), want)
	}
}

func TestTitleInfo_TitleName(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetTitleName("LONGTITLE")

	want := "+40LONGTITL"

	if ti.TitleName() != want {
		t.Errorf("TitleName = %s, wanted %s", ti.TitleName(), want)
	}
}

func TestTitleInfo_PublicationDate(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetPublicationDate("2020-07-06")

	want := "+95200706"

	if ti.PublicationDate() != want {
		t.Errorf("PublicationDate = %s, wanted %s", ti.PublicationDate(), want)
	}
}

func TestTitleInfo_CountryCode(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetCountryCode("13")

	want := "+9713"

	if ti.CountryCode() != want {
		t.Errorf("CountryCode = %s, wanted %s", ti.CountryCode(), want)
	}
}

func TestTitleInfo_PrintObjectColor(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetPrintObjectColor("12345678")

	want := "+9412345678"

	if ti.PrintObjectColor() != want {
		t.Errorf("PrintObjectColor = %s, wanted %s", ti.PrintObjectColor(), want)
	}
}

func TestTitleInfo_AdditionalInfo(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetAdditionalInfo("A really really long text with no real meaning!  ")

	want := "+08A really really long text with no real meaning!   "

	if ti.AdditionalInfo() != want {
		t.Errorf("AdditionalInfo = %s, wanted %s", ti.AdditionalInfo(), want)
	}
}

func TestTitleInfo_Complete(t *testing.T) {
	ti := NewTitleInfo()
	ti.SetPrintObjectName("FERAGTITLE")
	ti.SetTitleName("THETITLE")
	ti.SetPublicationDate("2020-07-06")
	ti.SetCountryCode("13")
	ti.SetPrintObjectColor("87654321")
	ti.SetAdditionalInfo("")

	want := "%2440+93FERAGTITLE  +40THETITLE+95200706+9713+9487654321!" + linebreak

	if ti.Message() != want {
		t.Errorf("Complete = |%s|", ti.Message())
		t.Errorf("  wanted = |%s|", want)
	}
}
