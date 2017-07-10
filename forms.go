// forms.go
package main

import (
    "github.com/votezilla/gforms"
)

// === FIELDS ===
var (
    username = gforms.NewTextField(
        "username",
        gforms.Validators{
            gforms.Required(),
            gforms.MaxLengthValidator(40),
        },
        gforms.TextInputWidget(map[string]string{
            "autocorrect": "off",
            "spellcheck": "false",
            "autocapitalize": "off",
            "autofocus": "true",
        }),
    )
    email = gforms.NewTextField(
        "email",
        gforms.Validators{
			gforms.EmailValidator(),
            gforms.Required(),
            gforms.MaxLengthValidator(345),          
        },
        gforms.TextInputWidget(map[string]string{
            "autocorrect": "off",
            "spellcheck": "false",
            "autocapitalize": "off",
        }),
    )    
    password = gforms.NewTextField(
        "password",
        gforms.Validators{
            gforms.Required(),
            gforms.MinLengthValidator(8),
            gforms.MaxLengthValidator(40),
            gforms.PasswordStrengthValidator(3), // Require strong password.
        },
    )
    confirmPassword = gforms.NewTextField(
        "confirm password",
        gforms.Validators{
            gforms.FieldMatchValidator("password"),
        },
        gforms.PasswordInputWidget(map[string]string{}),
    )
    rememberMe = gforms.NewTextField(
        "remember me",
        gforms.Validators{},
        gforms.CheckboxMultipleWidget(
            map[string]string{},
            func() gforms.CheckboxOptions { return gforms.StringCheckboxOptions([][]string{
                {"Stay logged in (uncheck if a shared computer)", "R", "false", "false"},
            })},
        ),
    )    
    
    // Demographics
    name = gforms.NewTextField(
        "full name",
        gforms.Validators{
            gforms.Required(),
            gforms.MaxLengthValidator(100),
            gforms.FullNameValidator(),
        },
        gforms.TextInputWidget(map[string]string{
            "autocorrect": "off",
            "spellcheck": "false",
        }),
    )
    birthYear = gforms.NewTextField( //NewFloatField(
        "year of birth",
        gforms.Validators{
            gforms.Required(),
            gforms.MinLengthValidator(4),
            gforms.MaxLengthValidator(4),
    })
    country = gforms.NewTextField(
        "country",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.SelectWidget(
            map[string]string{},
            func() gforms.SelectOptions { return gforms.StringSelectOptions([][]string{
                // ISO-3166-1 Alpha-2 country list.  See: https://www.freeformatter.com/iso-country-list-html-select.html
                {"-", "-", "false", "false"},
                {"United States", "US", "false", "false"},
                {"United States Minor Outlying Islands", "UM", "false", "false"},
                {"Afghanistan", "AF", "false", "false"},
                {"Åland Islands", "AX", "false", "false"},
                {"Albania", "AL", "false", "false"},
                {"Algeria", "DZ", "false", "false"},
                {"American Samoa", "AS", "false", "false"},
                {"Andorra", "AD", "false", "false"},
                {"Angola", "AO", "false", "false"},
                {"Anguilla", "AI", "false", "false"},
                {"Antarctica", "AQ", "false", "false"},
                {"Antigua and Barbuda", "AG", "false", "false"},
                {"Argentina", "AR", "false", "false"},
                {"Armenia", "AM", "false", "false"},
                {"Aruba", "AW", "false", "false"},
                {"Australia", "AU", "false", "false"},
                {"Austria", "AT", "false", "false"},
                {"Azerbaijan", "AZ", "false", "false"},
                {"Bahamas", "BS", "false", "false"},
                {"Bahrain", "BH", "false", "false"},
                {"Bangladesh", "BD", "false", "false"},
                {"Barbados", "BB", "false", "false"},
                {"Belarus", "BY", "false", "false"},
                {"Belgium", "BE", "false", "false"},
                {"Belize", "BZ", "false", "false"},
                {"Benin", "BJ", "false", "false"},
                {"Bermuda", "BM", "false", "false"},
                {"Bhutan", "BT", "false", "false"},
                {"Bolivia, Plurinational State of", "BO", "false", "false"},
                {"Bonaire, Sint Eustatius and Saba", "BQ", "false", "false"},
                {"Bosnia and Herzegovina", "BA", "false", "false"},
                {"Botswana", "BW", "false", "false"},
                {"Bouvet Island", "BV", "false", "false"},
                {"Brazil", "BR", "false", "false"},
                {"British Indian Ocean Territory", "IO", "false", "false"},
                {"Brunei Darussalam", "BN", "false", "false"},
                {"Bulgaria", "BG", "false", "false"},
                {"Burkina Faso", "BF", "false", "false"},
                {"Burundi", "BI", "false", "false"},
                {"Cambodia", "KH", "false", "false"},
                {"Cameroon", "CM", "false", "false"},
                {"Canada", "CA", "false", "false"},
                {"Cape Verde", "CV", "false", "false"},
                {"Cayman Islands", "KY", "false", "false"},
                {"Central African Republic", "CF", "false", "false"},
                {"Chad", "TD", "false", "false"},
                {"Chile", "CL", "false", "false"},
                {"China", "CN", "false", "false"},
                {"Christmas Island", "CX", "false", "false"},
                {"Cocos (Keeling) Islands", "CC", "false", "false"},
                {"Colombia", "CO", "false", "false"},
                {"Comoros", "KM", "false", "false"},
                {"Congo", "CG", "false", "false"},
                {"Congo, the Democratic Republic of the", "CD", "false", "false"},
                {"Cook Islands", "CK", "false", "false"},
                {"Costa Rica", "CR", "false", "false"},
                {"Côte d'Ivoire", "CI", "false", "false"},
                {"Croatia", "HR", "false", "false"},
                {"Cuba", "CU", "false", "false"},
                {"Curaçao", "CW", "false", "false"},
                {"Cyprus", "CY", "false", "false"},
                {"Czech Republic", "CZ", "false", "false"},
                {"Denmark", "DK", "false", "false"},
                {"Djibouti", "DJ", "false", "false"},
                {"Dominica", "DM", "false", "false"},
                {"Dominican Republic", "DO", "false", "false"},
                {"Ecuador", "EC", "false", "false"},
                {"Egypt", "EG", "false", "false"},
                {"El Salvador", "SV", "false", "false"},
                {"Equatorial Guinea", "GQ", "false", "false"},
                {"Eritrea", "ER", "false", "false"},
                {"Estonia", "EE", "false", "false"},
                {"Ethiopia", "ET", "false", "false"},
                {"Falkland Islands (Malvinas)", "FK", "false", "false"},
                {"Faroe Islands", "FO", "false", "false"},
                {"Fiji", "FJ", "false", "false"},
                {"Finland", "FI", "false", "false"},
                {"France", "FR", "false", "false"},
                {"French Guiana", "GF", "false", "false"},
                {"French Polynesia", "PF", "false", "false"},
                {"French Southern Territories", "TF", "false", "false"},
                {"Gabon", "GA", "false", "false"},
                {"Gambia", "GM", "false", "false"},
                {"Georgia", "GE", "false", "false"},
                {"Germany", "DE", "false", "false"},
                {"Ghana", "GH", "false", "false"},
                {"Gibraltar", "GI", "false", "false"},
                {"Greece", "GR", "false", "false"},
                {"Greenland", "GL", "false", "false"},
                {"Grenada", "GD", "false", "false"},
                {"Guadeloupe", "GP", "false", "false"},
                {"Guam", "GU", "false", "false"},
                {"Guatemala", "GT", "false", "false"},
                {"Guernsey", "GG", "false", "false"},
                {"Guinea", "GN", "false", "false"},
                {"Guinea-Bissau", "GW", "false", "false"},
                {"Guyana", "GY", "false", "false"},
                {"Haiti", "HT", "false", "false"},
                {"Heard Island and McDonald Islands", "HM", "false", "false"},
                {"Holy See (Vatican City State)", "VA", "false", "false"},
                {"Honduras", "HN", "false", "false"},
                {"Hong Kong", "HK", "false", "false"},
                {"Hungary", "HU", "false", "false"},
                {"Iceland", "IS", "false", "false"},
                {"India", "IN", "false", "false"},
                {"Indonesia", "ID", "false", "false"},
                {"Iran, Islamic Republic of", "IR", "false", "false"},
                {"Iraq", "IQ", "false", "false"},
                {"Ireland", "IE", "false", "false"},
                {"Isle of Man", "IM", "false", "false"},
                {"Israel", "IL", "false", "false"},
                {"Italy", "IT", "false", "false"},
                {"Jamaica", "JM", "false", "false"},
                {"Japan", "JP", "false", "false"},
                {"Jersey", "JE", "false", "false"},
                {"Jordan", "JO", "false", "false"},
                {"Kazakhstan", "KZ", "false", "false"},
                {"Kenya", "KE", "false", "false"},
                {"Kiribati", "KI", "false", "false"},
                {"Korea, Democratic People's Republic of", "KP", "false", "false"},
                {"Korea, Republic of", "KR", "false", "false"},
                {"Kuwait", "KW", "false", "false"},
                {"Kyrgyzstan", "KG", "false", "false"},
                {"Lao People's Democratic Republic", "LA", "false", "false"},
                {"Latvia", "LV", "false", "false"},
                {"Lebanon", "LB", "false", "false"},
                {"Lesotho", "LS", "false", "false"},
                {"Liberia", "LR", "false", "false"},
                {"Libya", "LY", "false", "false"},
                {"Liechtenstein", "LI", "false", "false"},
                {"Lithuania", "LT", "false", "false"},
                {"Luxembourg", "LU", "false", "false"},
                {"Macao", "MO", "false", "false"},
                {"Macedonia, the former Yugoslav Republic of", "MK", "false", "false"},
                {"Madagascar", "MG", "false", "false"},
                {"Malawi", "MW", "false", "false"},
                {"Malaysia", "MY", "false", "false"},
                {"Maldives", "MV", "false", "false"},
                {"Mali", "ML", "false", "false"},
                {"Malta", "MT", "false", "false"},
                {"Marshall Islands", "MH", "false", "false"},
                {"Martinique", "MQ", "false", "false"},
                {"Mauritania", "MR", "false", "false"},
                {"Mauritius", "MU", "false", "false"},
                {"Mayotte", "YT", "false", "false"},
                {"Mexico", "MX", "false", "false"},
                {"Micronesia, Federated States of", "FM", "false", "false"},
                {"Moldova, Republic of", "MD", "false", "false"},
                {"Monaco", "MC", "false", "false"},
                {"Mongolia", "MN", "false", "false"},
                {"Montenegro", "ME", "false", "false"},
                {"Montserrat", "MS", "false", "false"},
                {"Morocco", "MA", "false", "false"},
                {"Mozambique", "MZ", "false", "false"},
                {"Myanmar", "MM", "false", "false"},
                {"Namibia", "NA", "false", "false"},
                {"Nauru", "NR", "false", "false"},
                {"Nepal", "NP", "false", "false"},
                {"Netherlands", "NL", "false", "false"},
                {"New Caledonia", "NC", "false", "false"},
                {"New Zealand", "NZ", "false", "false"},
                {"Nicaragua", "NI", "false", "false"},
                {"Niger", "NE", "false", "false"},
                {"Nigeria", "NG", "false", "false"},
                {"Niue", "NU", "false", "false"},
                {"Norfolk Island", "NF", "false", "false"},
                {"Northern Mariana Islands", "MP", "false", "false"},
                {"Norway", "NO", "false", "false"},
                {"Oman", "OM", "false", "false"},
                {"Pakistan", "PK", "false", "false"},
                {"Palau", "PW", "false", "false"},
                {"Palestinian Territory, Occupied", "PS", "false", "false"},
                {"Panama", "PA", "false", "false"},
                {"Papua New Guinea", "PG", "false", "false"},
                {"Paraguay", "PY", "false", "false"},
                {"Peru", "PE", "false", "false"},
                {"Philippines", "PH", "false", "false"},
                {"Pitcairn", "PN", "false", "false"},
                {"Poland", "PL", "false", "false"},
                {"Portugal", "PT", "false", "false"},
                {"Puerto Rico", "PR", "false", "false"},
                {"Qatar", "QA", "false", "false"},
                {"Réunion", "RE", "false", "false"},
                {"Romania", "RO", "false", "false"},
                {"Russia", "RU", "false", "false"},
                {"Rwanda", "RW", "false", "false"},
                {"Saint Barthélemy", "BL", "false", "false"},
                {"Saint Helena, Ascension and Tristan da Cunha", "SH", "false", "false"},
                {"Saint Kitts and Nevis", "KN", "false", "false"},
                {"Saint Lucia", "LC", "false", "false"},
                {"Saint Martin (French part)", "MF", "false", "false"},
                {"Saint Pierre and Miquelon", "PM", "false", "false"},
                {"Saint Vincent and the Grenadines", "VC", "false", "false"},
                {"Samoa", "WS", "false", "false"},
                {"San Marino", "SM", "false", "false"},
                {"Sao Tome and Principe", "ST", "false", "false"},
                {"Saudi Arabia", "SA", "false", "false"},
                {"Senegal", "SN", "false", "false"},
                {"Serbia", "RS", "false", "false"},
                {"Seychelles", "SC", "false", "false"},
                {"Sierra Leone", "SL", "false", "false"},
                {"Singapore", "SG", "false", "false"},
                {"Sint Maarten (Dutch part)", "SX", "false", "false"},
                {"Slovakia", "SK", "false", "false"},
                {"Slovenia", "SI", "false", "false"},
                {"Solomon Islands", "SB", "false", "false"},
                {"Somalia", "SO", "false", "false"},
                {"South Africa", "ZA", "false", "false"},
                {"South Georgia and the South Sandwich Islands", "GS", "false", "false"},
                {"South Sudan", "SS", "false", "false"},
                {"Spain", "ES", "false", "false"},
                {"Sri Lanka", "LK", "false", "false"},
                {"Sudan", "SD", "false", "false"},
                {"Suriname", "SR", "false", "false"},
                {"Svalbard and Jan Mayen", "SJ", "false", "false"},
                {"Swaziland", "SZ", "false", "false"},
                {"Sweden", "SE", "false", "false"},
                {"Switzerland", "CH", "false", "false"},
                {"Syrian Arab Republic", "SY", "false", "false"},
                {"Taiwan, Province of China", "TW", "false", "false"},
                {"Tajikistan", "TJ", "false", "false"},
                {"Tanzania, United Republic of", "TZ", "false", "false"},
                {"Thailand", "TH", "false", "false"},
                {"Timor-Leste", "TL", "false", "false"},
                {"Togo", "TG", "false", "false"},
                {"Tokelau", "TK", "false", "false"},
                {"Tonga", "TO", "false", "false"},
                {"Trinidad and Tobago", "TT", "false", "false"},
                {"Tunisia", "TN", "false", "false"},
                {"Turkey", "TR", "false", "false"},
                {"Turkmenistan", "TM", "false", "false"},
                {"Turks and Caicos Islands", "TC", "false", "false"},
                {"Tuvalu", "TV", "false", "false"},
                {"Uganda", "UG", "false", "false"},
                {"Ukraine", "UA", "false", "false"},
                {"United Arab Emirates", "AE", "false", "false"},
                {"United Kingdom", "GB", "false", "false"},
                {"United States", "US", "false", "false"},
                {"United States Minor Outlying Islands", "UM", "false", "false"},
                {"Uruguay", "UY", "false", "false"},
                {"Uzbekistan", "UZ", "false", "false"},
                {"Vanuatu", "VU", "false", "false"},
                {"Venezuela", "VE", "false", "false"},
                {"Viet Nam", "VN", "false", "false"},
                {"Virgin Islands, British", "VG", "false", "false"},
                {"Virgin Islands, U.S.", "VI", "false", "false"},
                {"Wallis and Futuna", "WF", "false", "false"},
                {"Western Sahara", "EH", "false", "false"},
                {"Yemen", "YE", "false", "false"},
                {"Zambia", "ZM", "false", "false"},
                {"Zimbabwe", "ZW", "false", "false"},
            })},
        ),
    )
    location = gforms.NewTextField(
        "location",
        gforms.Validators{
            gforms.Required(),
            gforms.MinLengthValidator(5),
            gforms.MaxLengthValidator(60),
    })
    gender = gforms.NewTextField(
        "gender",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.SelectWidget(
            map[string]string{},
            func() gforms.SelectOptions { return gforms.StringSelectOptions([][]string{
                {"-",      "-", "false", "false"},
                {"Male",   "M", "false", "false"},
                {"Female", "F", "false", "false"},
                {"Other",  "O", "false", "false"},
            })},
        ),
    )
    party = gforms.NewTextField(
        "party",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.SelectWidget(
            map[string]string{},
            func() gforms.SelectOptions { return gforms.StringSelectOptions([][]string{
                {"-",           "-", "false", "false"},
                {"Republican",  "R", "false", "false"},
                {"Democrat",    "D", "false", "false"},
                {"Independent", "I", "false", "false"},
                {"Other",       "O", "false", "false"},
            })},
        ),
    )
    race = gforms.NewMultipleTextField(
		"race / ethnicity",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.CheckboxMultipleWidget(
            map[string]string{},
            func() gforms.CheckboxOptions { return gforms.StringCheckboxOptions([][]string{
                {"American Indian or Alaska Native",          "I", "false", "false"},
                {"Asian",                                     "A", "false", "false"},
                {"Black or African American",                 "B", "false", "false"},
                {"Hispanic, Latino, or Spanish",              "H", "false", "false"},
                {"Native Hawaiian or Other Pacific Islander", "P", "false", "false"},
                {"White",                                     "W", "false", "false"},
                {"Other",                                     "O", "false", "false"},
            })},
        ),
    )
    marital = gforms.NewTextField(
        "marital status",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.SelectWidget(
            map[string]string{},
            func() gforms.SelectOptions { return gforms.StringSelectOptions([][]string{
                {"-",                               "-", "false", "false"},
                {"Single (Never Married)",          "S", "false", "false"},
                {"Divorced or Separated",           "D", "false", "false"},
                {"Widowed",                         "W", "false", "false"},
                {"Married or Domestic Partnership", "M", "false", "false"},
            })},
        ),
    )   
    schooling = gforms.NewTextField(
        "furthest schooling completed",
        gforms.Validators{
            gforms.Required(),
        },
        gforms.SelectWidget(
            map[string]string{},
            func() gforms.SelectOptions { return gforms.StringSelectOptions([][]string{
                {"-",                                "-", "false", "false"},
                {"Less than a high school diploma",  "L", "false", "false"},
                {"High school degree or equivalent", "H", "false", "false"},
                {"Some college, but no degree",      "S", "false", "false"},
                {"College graduate",                 "C", "false", "false"},
                {"Postgraduate study",               "P", "false", "false"},
            })},
        ),
    )
)

// === COUNTRY DATA ===
var (
	// https://en.m.wikipedia.org/wiki/Federated_state
	CountriesWithStates = map[string]bool{"AE":true,"AR":true,"AT":true,"AU":true,"BA":true,"BE":true,"BR":true,"CA":true,"CH":true,"DE":true,"ET":true,"FM":true,"IN":true,"IQ":true,"KM":true,"KN":true,"MX":true,"MY":true,"NG":true,"NP":true,"PK":true,"RU":true,"SD":true,"SO":true,"SS":true,"US":true,"VE":true}
	
	// https://www.ups.com/worldshiphelp/WS16/ENU/AppHelp/Codes/Countries_Territories_Requiring_Postal_Codes.htm
	CountriesWithPostalCodes = map[string]bool{"A2":true,"AM":true,"AR":true,"AT":true,"AU":true,"AZ":true,"BA":true,"BD":true,"BE":true,"BG":true,"BN":true,"BR":true,"BY":true,"CA":true,"CH":true,"CN":true,"CS":true,"CY":true,"CZ":true,"DE":true,"DK":true,"DZ":true,"EE":true,"EN":true,"ES":true,"FI":true,"FO":true,"FR":true,"GB":true,"GE":true,"GG":true,"GL":true,"GR":true,"GU":true,"HO":true,"HR":true,"HU":true,"IC":true,"ID":true,"IL":true,"IN":true,"IT":true,"JE":true,"JP":true,"KG":true,"KO":true,"KR":true,"KZ":true,"LI":true,"LK":true,"LT":true,"LU":true,"LV":true,"M3":true,"ME":true,"MG":true,"MH":true,"MK":true,"MN":true,"MQ":true,"MX":true,"MY":true,"NB":true,"NL":true,"NO":true,"NT":true,"NZ":true,"PH":true,"PK":true,"PL":true,"PO":true,"PR":true,"PT":true,"RE":true,"RU":true,"SA":true,"SE":true,"SF":true,"SG":true,"SI":true,"SK":true,"SX":true,"TH":true,"TJ":true,"TM":true,"TN":true,"TR":true,"TU":true,"TW":true,"UA":true,"US":true,"UV":true,"UY":true,"UZ":true,"VA":true,"VI":true,"VL":true,"VN":true,"WL":true,"YA":true,"YT":true,"ZA":true}
)

// === FORM POST DATA ===
type LoginData struct {
    Email                   string `gforms:"email"`
    Password                string `gforms:"password"`
    RememberMe              string `gforms:"remember me"`
}

type RegisterData struct {
    Email                	string `gforms:"email"`
    Password                string `gforms:"password"`
}

type RegisterDetailsData struct {
    Name                    string `gforms:"full name"`

    // location
    Country                 string `gforms:"country"`
    Location                string `gforms:"location"`

    // demographic
    BirthYear               string `gforms:"year of birth"`
    Gender                  string `gforms:"gender"`
    Party                   string `gforms:"party"`
    Races					[]string `gforms:"race / ethnicity"`
    Marital                 string `gforms:"marital status"`
    Schooling               string `gforms:"furthest schooling completed"`
}

// === FORMS ===
var (
    LoginForm = gforms.DefineForm(gforms.NewFields(
        email,
        password,
        rememberMe,
    ))
    RegisterForm = gforms.DefineForm(gforms.NewFields(
        email,
        password,
    ))
    RegisterDetailsForm = gforms.DefineForm(gforms.NewFields(
        // name
        name,
        
        // location
        country,
        location,

        // demographic
        birthYear,
        gender,
        party,
        race,
        marital,
        schooling,      
    ))
) // var

// === FORM TYPES ===
type TableForm struct {
	Form			*gforms.FormInstance
	CallToAction	string
	AdditionalError string
}

type FormArgs struct{
	Forms			[]TableForm
	Title			string
	Introduction	string
	Footer			string
	Script			string
}