package gootp

import "testing"

var MD4_ENCODINGS = []struct {
	passphrase string
	seed       string
	seq        int
	otp        string
}{
	{"This is a test.", "TeSt", 0, "ROME MUG FRED SCAN LIVE LACE"},
	{"This is a test.", "TeSt", 1, "CARD SAD MINI RYE COL KIN"},
	{"This is a test.", "TeSt", 99, "NOTE OUT IBIS SINK NAVE MODE"},
	{"AbCdEfGhIjK", "alpha1", 0, "AWAY SEN ROOK SALT LICE MAP"},
	{"AbCdEfGhIjK", "alpha1", 1, "CHEW GRIM WU HANG BUCK SAID"},
	{"AbCdEfGhIjK", "alpha1", 99, "ROIL FREE COG HUNK WAIT COCA"},
	{"OTP's are good", "correct", 0, "FOOL STEM DONE TOOL BECK NILE"},
	{"OTP's are good", "correct", 1, "GIST AMOS MOOT AIDS FOOD SEEM"},
	{"OTP's are good", "correct", 99, "TAG SLOW NOV MIN WOOL KENO"},
}

func TestOtpMd4(t *testing.T) {
	for _, encoding := range MD4_ENCODINGS {
		otp := OtpMd4(encoding.seq, encoding.seed, encoding.passphrase)
		if otp != encoding.otp {
			t.Errorf("Unexpected output: %s for input: %s", otp, encoding)
		}
	}
}

var MD5_ENCODINGS = []struct {
	passphrase string
	seed       string
	seq        int
	otp        string
}{
	{"This is a test.", "TeSt", 0, "INCH SEA ANNE LONG AHEM TOUR"},
	{"This is a test.", "TeSt", 1, "EASE OIL FUM CURE AWRY AVIS"},
	{"This is a test.", "TeSt", 99, "BAIL TUFT BITS GANG CHEF THY"},
	{"AbCdEfGhIjK", "alpha1", 0, "FULL PEW DOWN ONCE MORT ARC"},
	{"AbCdEfGhIjK", "alpha1", 1, "FACT HOOF AT FIST SITE KENT"},
	{"AbCdEfGhIjK", "alpha1", 99, "BODE HOP JAKE STOW JUT RAP"},
	{"OTP's are good", "correct", 0, "ULAN NEW ARMY FUSE SUIT EYED"},
	{"OTP's are good", "correct", 1, "SKIM CULT LOB SLAM POE HOWL"},
	{"OTP's are good", "correct", 99, "LONG IVY JULY AJAR BOND LEE"},
}

func TestOtpMd5(t *testing.T) {
	for _, encoding := range MD5_ENCODINGS {
		otp := OtpMd5(encoding.seq, encoding.seed, encoding.passphrase)
		if otp != encoding.otp {
			t.Errorf("Unexpected output: %s for input: %s", otp, encoding)
		}
	}
}

var SHA1_ENCODINGS = []struct {
	passphrase string
	seed       string
	seq        int
	otp        string
}{
	{"This is a test.", "TeSt", 0, "MILT VARY MAST OK SEES WENT"},
	{"This is a test.", "TeSt", 1, "CART OTTO HIVE ODE VAT NUT"},
	{"This is a test.", "TeSt", 99, "GAFF WAIT SKID GIG SKY EYED"},
	{"AbCdEfGhIjK", "alpha1", 0, "LEST OR HEEL SCOT ROB SUIT"},
	{"AbCdEfGhIjK", "alpha1", 1, "RITE TAKE GELD COST TUNE RECK"},
	{"AbCdEfGhIjK", "alpha1", 99, "MAY STAR TIN LYON VEDA STAN"},
	{"OTP's are good", "correct", 0, "RUST WELT KICK FELL TAIL FRAU"},
	{"OTP's are good", "correct", 1, "FLIT DOSE ALSO MEW DRUM DEFY"},
	{"OTP's are good", "correct", 99, "AURA ALOE HURL WING BERG WAIT"},
}

func TestOtpSha1(t *testing.T) {
	for _, encoding := range SHA1_ENCODINGS {
		otp := OtpSha1(encoding.seq, encoding.seed, encoding.passphrase)
		if otp != encoding.otp {
			t.Errorf("Unexpected output: %s for input: %s", otp, encoding)
		}
	}
}
