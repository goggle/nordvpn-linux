package internal

import (
	"testing"

	"github.com/NordSecurity/nordvpn-linux/test/category"

	"github.com/stretchr/testify/assert"
)

func TestExtractZone(t *testing.T) {
	category.Set(t, category.Unit)

	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `Timezone=Europe/Vilnius
LocalRTC=no
CanNTP=yes
NTP=yes
NTPSynchronized=yes
TimeUSec=Wed 2020-05-06 17:06:01 EEST
RTCTimeUSec=Wed 2020-05-06 17:06:01 EEST
`,
			expected: "Europe/Vilnius",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		zone := extractZone([]byte(tt.input))
		assert.Equal(t, tt.expected, zone)
	}
}

func TestObfuscate(t *testing.T) {
	category.Set(t, category.Unit)

	tests := []struct {
		timezone string
		hash     string
	}{
		// Africa
		{"Africa/Abidjan", "8e9e0533acc53b"},
		{"Africa/Accra", "413c61b75c1786"},
		{"Africa/Addis_Ababa", "441ada017217a9"},
		{"Africa/Algiers", "a011131551bd4e"},
		{"Africa/Asmara", "8b802208da0d54"},
		{"Africa/Bamako", "c7c63329217f6a"},
		{"Africa/Bangui", "886e7fd3eee8ac"},
		{"Africa/Banjul", "fbeba3a5ec59fd"},
		{"Africa/Bissau", "467771c29cd75a"},
		{"Africa/Blantyre", "6d19c74920d66b"},
		{"Africa/Brazzaville", "e73e6efe2a81a2"},
		{"Africa/Bujumbura", "45ceddf2f828ed"},
		{"Africa/Cairo", "2e9a9bb3cc4394"},
		{"Africa/Casablanca", "9824257c811fd5"},
		{"Africa/Ceuta", "25e1c83126001a"},
		{"Africa/Conakry", "573b6836c1f9dd"},
		{"Africa/Dakar", "47302d2c6d23bf"},
		{"Africa/Dar_es_Salaam", "8bedeb39bc6db4"},
		{"Africa/Djibouti", "9da4c7b429018a"},
		{"Africa/Douala", "68e4f331b314c8"},
		{"Africa/El_Aaiun", "39923d61d9617e"},
		{"Africa/Freetown", "957ec0de24ed1c"},
		{"Africa/Gaborone", "bf7846ff85a6fe"},
		{"Africa/Harare", "ac504bd693c006"},
		{"Africa/Johannesburg", "cc7db809992d5b"},
		{"Africa/Juba", "3b62a862271521"},
		{"Africa/Kampala", "84634579fea52f"},
		{"Africa/Khartoum", "e9e3b0177711e8"},
		{"Africa/Kigali", "f32e5b5922d822"},
		{"Africa/Kinshasa", "94e722c4a361d9"},
		{"Africa/Lagos", "93ba6a13255301"},
		{"Africa/Libreville", "8fd743d7caf53d"},
		{"Africa/Lome", "788aab353fccfb"},
		{"Africa/Luanda", "448b74968f8095"},
		{"Africa/Lubumbashi", "1b902fa1eb2f2d"},
		{"Africa/Lusaka", "e7cc2eb182c019"},
		{"Africa/Malabo", "ec05bc3133198e"},
		{"Africa/Maputo", "8050abe3fb4ba7"},
		{"Africa/Maseru", "baca8339847c1d"},
		{"Africa/Mbabane", "90e6efc1953e44"},
		{"Africa/Mogadishu", "e8dc340bb2411c"},
		{"Africa/Monrovia", "c27f4f2d413056"},
		{"Africa/Nairobi", "8a46a5b88726e1"},
		{"Africa/Ndjamena", "53dbf03db9e7ae"},
		{"Africa/Niamey", "4c4c3af4718e88"},
		{"Africa/Nouakchott", "f94c4a85f18952"},
		{"Africa/Ouagadougou", "85ebcfb3d38156"},
		{"Africa/Porto-Novo", "b4677dd13ead6e"},
		{"Africa/Sao_Tome", "921e22e0c2264f"},
		{"Africa/Tripoli", "77e2718712382d"},
		{"Africa/Tunis", "6dbca29ce898b3"},
		{"Africa/Windhoek", "507df7b946afca"},

		// America
		{"America/Adak", "3e1e07e40908ba"},
		{"America/Anchorage", "d6bd4d6652b20d"},
		{"America/Anguilla", "4f89952748d92b"},
		{"America/Antigua", "0300c9b6089b95"},
		{"America/Araguaina", "812dea334a55c4"},
		{"America/Argentina/Buenos_Aires", "2d107ae4da95ed"},
		{"America/Argentina/Catamarca", "4e2bffac54b369"},
		{"America/Argentina/Cordoba", "8575692fa08ef1"},
		{"America/Argentina/Jujuy", "5f7e248ce89364"},
		{"America/Argentina/La_Rioja", "9bbec736626a75"},
		{"America/Argentina/Mendoza", "5a746445e01b78"},
		{"America/Argentina/Rio_Gallegos", "49e12d39317745"},
		{"America/Argentina/Salta", "25ed1c53b630ea"},
		{"America/Argentina/San_Juan", "3feda1eb090939"},
		{"America/Argentina/San_Luis", "44c9bd0297f298"},
		{"America/Argentina/Tucuman", "908883d0e43f66"},
		{"America/Argentina/Ushuaia", "9a6ca963dc13ae"},
		{"America/Aruba", "c8d9bcf7660d0b"},
		{"America/Asuncion", "8ca340d8332bd5"},
		{"America/Atikokan", "aa6a796f3a7830"},
		{"America/Bahia", "4bcf2327bdfadd"},
		{"America/Bahia_Banderas", "27414574ddea32"},
		{"America/Barbados", "f985cbcbd04ee1"},
		{"America/Belem", "edf73aebb0509f"},
		{"America/Belize", "24381ca5e5845a"},
		{"America/Blanc-Sablon", "0c61ac7b3ee6a0"},
		{"America/Boa_Vista", "255f16ab6c19ac"},
		{"America/Bogota", "43a0373f2cad29"},
		{"America/Boise", "4d8533a4cf39e8"},
		{"America/Cambridge_Bay", "e6383dcfed6a3a"},
		{"America/Campo_Grande", "8f550bbe67456d"},
		{"America/Cancun", "d146e1f1fc8812"},
		{"America/Caracas", "7eb53a2fe44ecf"},
		{"America/Cayenne", "cc347c64f74da7"},
		{"America/Cayman", "75780bdb313a93"},
		{"America/Chicago", "a77b7136d7ecc3"},
		{"America/Chihuahua", "8c39ea6b0ab46b"},
		{"America/Costa_Rica", "4fddfbc439b7ee"},
		{"America/Creston", "bf88544b926f1b"},
		{"America/Cuiaba", "5ebaa1ca743cd3"},
		{"America/Curacao", "0166b5e2998408"},
		{"America/Danmarkshavn", "d1a448a62d1f64"},
		{"America/Dawson", "650a86e5095b91"},
		{"America/Dawson_Creek", "12d02d2d9a13b5"},
		{"America/Denver", "9e36365b2f6cfe"},
		{"America/Detroit", "034883cbbcdc28"},
		{"America/Dominica", "6d6eade7ca6ca9"},
		{"America/Edmonton", "58f2f20328fcc9"},
		{"America/Eirunepe", "6ac36e16b23179"},
		{"America/El_Salvador", "40e46ad3f9f21e"},
		{"America/Fort_Nelson", "3a65ddfba76732"},
		{"America/Fortaleza", "1dc70db68b2bcf"},
		{"America/Glace_Bay", "b4c48e5a8ca124"},
		{"America/Godthab", "5e8c06da3febfc"},
		{"America/Goose_Bay", "e5b7d173692bcf"},
		{"America/Grand_Turk", "b10832f70be435"},
		{"America/Grenada", "fa99bd35c952d8"},
		{"America/Guadeloupe", "8f7b5f1a5bb27a"},
		{"America/Guatemala", "83e28b1c7cce54"},
		{"America/Guayaquil", "2bb9c2965be336"},
		{"America/Guyana", "94e1ec24f64af8"},
		{"America/Halifax", "9958a325c3e38a"},
		{"America/Havana", "0093ef77adba1a"},
		{"America/Hermosillo", "7bd503343f8307"},
		{"America/Indiana/Indianapolis", "983fc329961bf2"},
		{"America/Indiana/Knox", "a7c2845512f2e1"},
		{"America/Indiana/Marengo", "305c1736bbb9d6"},
		{"America/Indiana/Petersburg", "11e160f8fc0101"},
		{"America/Indiana/Tell_City", "32a6efd42db8a5"},
		{"America/Indiana/Vevay", "9a85937e48c4bc"},
		{"America/Indiana/Vincennes", "1d3d116d624bf5"},
		{"America/Indiana/Winamac", "42867a5628737d"},
		{"America/Inuvik", "51299c7c76af25"},
		{"America/Iqaluit", "a841c860c9029d"},
		{"America/Jamaica", "390566b0d523bf"},
		{"America/Juneau", "64a28b6ab0f855"},
		{"America/Kentucky/Louisville", "c36bcbab2ba4e1"},
		{"America/Kentucky/Monticello", "31671ebd2d168d"},
		{"America/Kralendijk", "6fd0ae0167c93a"},
		{"America/La_Paz", "c5680a19e2f96b"},
		{"America/Lima", "d4914fc8e52f2b"},
		{"America/Los_Angeles", "2d4bbedff80785"},
		{"America/Lower_Princes", "e6bfe941e700e9"},
		{"America/Maceio", "2b4b52ed6662f7"},
		{"America/Managua", "cab62ce0058178"},
		{"America/Manaus", "e42c8b46eba659"},
		{"America/Marigot", "ad2ca32e3871b1"},
		{"America/Martinique", "7d07c17a6465dc"},
		{"America/Matamoros", "1535ac2efbde38"},
		{"America/Mazatlan", "021bd35a0bd275"},
		{"America/Menominee", "65e393f664b176"},
		{"America/Merida", "5df6de9cb5ace7"},
		{"America/Metlakatla", "e6e6dae1f34ffc"},
		{"America/Mexico_City", "fd8d3b93c3d8ae"},
		{"America/Miquelon", "0717decf4459cb"},
		{"America/Moncton", "b3215ec38257b7"},
		{"America/Monterrey", "4e6c351c76ec5c"},
		{"America/Montevideo", "51ffc663078e20"},
		{"America/Montserrat", "1664fbe26abbe4"},
		{"America/Nassau", "ee0edae093f3a0"},
		{"America/New_York", "acb2a3287bc63b"},
		{"America/Nipigon", "37bd52f27fcc85"},
		{"America/Nome", "5d8d80c8d91cb7"},
		{"America/Noronha", "5e83381a8e4964"},
		{"America/North_Dakota/Beulah", "044bb9faa92e69"},
		{"America/North_Dakota/Center", "ba741b26bdbfe6"},
		{"America/North_Dakota/New_Salem", "603cceed36985d"},
		{"America/Ojinaga", "a6bec8fb1b3771"},
		{"America/Panama", "5008245115c3c1"},
		{"America/Pangnirtung", "6da96966b40259"},
		{"America/Paramaribo", "5a046b252eab60"},
		{"America/Phoenix", "b919121b45aeca"},
		{"America/Port-au-Prince", "165528b666cdbc"},
		{"America/Port_of_Spain", "d389d35d9924b9"},
		{"America/Porto_Velho", "a1e5ce5a6fa2e3"},
		{"America/Puerto_Rico", "c8e719bdfece90"},
		{"America/Punta_Arenas", "bf893bd53abaed"},
		{"America/Rainy_River", "056122f61e643f"},
		{"America/Rankin_Inlet", "ffb2a4357131ec"},
		{"America/Recife", "7dd4d6f32ecd57"},
		{"America/Regina", "24c66d1bae70c6"},
		{"America/Resolute", "a2d15c87291f18"},
		{"America/Rio_Branco", "e3b87ba408d1cc"},
		{"America/Santarem", "a7581d3c3f6741"},
		{"America/Santiago", "58c885a79d9850"},
		{"America/Santo_Domingo", "f1c0a3eb2c293f"},
		{"America/Sao_Paulo", "53e24af00e9630"},
		{"America/Scoresbysund", "4c5d2fd44a2b12"},
		{"America/Sitka", "a62ed7605f1f85"},
		{"America/St_Barthelemy", "d6b47c700f1da1"},
		{"America/St_Johns", "7aae522a697ecf"},
		{"America/St_Kitts", "08a0e68352ca3a"},
		{"America/St_Lucia", "70a9ef58820b35"},
		{"America/St_Thomas", "88b067a46e3577"},
		{"America/St_Vincent", "27838240d16089"},
		{"America/Swift_Current", "d56db9f453ad49"},
		{"America/Tegucigalpa", "da21eb2c2af548"},
		{"America/Thule", "e9ff957d161f46"},
		{"America/Thunder_Bay", "7721e607ac3959"},
		{"America/Tijuana", "73a0af6b24caae"},
		{"America/Toronto", "feaa40fb2ab0c6"},
		{"America/Tortola", "690111cb72de62"},
		{"America/Vancouver", "9e906781f47b12"},
		{"America/Whitehorse", "1813c1b86f1a72"},
		{"America/Winnipeg", "05551849429fb3"},
		{"America/Yakutat", "90f715c40f64cd"},
		{"America/Yellowknife", "e718ec0e274459"},

		// Antarctica
		{"Antarctica/Casey", "0c4b851a6bcddb"},
		{"Antarctica/Davis", "871eeaa9d9f759"},
		{"Antarctica/DumontDUrville", "e5a2f75d2ae842"},
		{"Antarctica/Macquarie", "f4296b07746616"},
		{"Antarctica/Mawson", "c36592a27bd50e"},
		{"Antarctica/McMurdo", "bd1de2d0b576e7"},
		{"Antarctica/Palmer", "8d00e16a35ea9d"},
		{"Antarctica/Rothera", "06aa1cba95d064"},
		{"Antarctica/Syowa", "b2d49028a72a99"},
		{"Antarctica/Troll", "0c9c095685c4fd"},
		{"Antarctica/Vostok", "76bc8cbfd5cf2c"},

		// Arctic
		{"Arctic/Longyearbyen", "54b2420865695b"},

		// Asia
		{"Asia/Aden", "49b11395442211"},
		{"Asia/Almaty", "7bb2ce5e75a6e0"},
		{"Asia/Amman", "07e1927c3b8aa0"},
		{"Asia/Anadyr", "add02c73b19bce"},
		{"Asia/Aqtau", "d247a085ce3442"},
		{"Asia/Aqtobe", "1a92c723b32270"},
		{"Asia/Ashgabat", "8d3bcac67cfee2"},
		{"Asia/Atyrau", "838f380d4b8bfc"},
		{"Asia/Baghdad", "cc688922b81879"},
		{"Asia/Bahrain", "bef0674680dc17"},
		{"Asia/Baku", "9818606cf4e6b9"},
		{"Asia/Bangkok", "e8def32656206d"},
		{"Asia/Barnaul", "e20e3f9d1d81c5"},
		{"Asia/Beirut", "47d69d9dd207a1"},
		{"Asia/Bishkek", "5d78121fd7289a"},
		{"Asia/Brunei", "0c0ef577e60c0e"},
		{"Asia/Chita", "a241cf7f63704f"},
		{"Asia/Choibalsan", "26813a92224fa6"},
		{"Asia/Colombo", "aa94105757d7c7"},
		{"Asia/Damascus", "dcc4971cdc69c5"},
		{"Asia/Dhaka", "d2342499eace38"},
		{"Asia/Dili", "75316c5b67abb6"},
		{"Asia/Dubai", "0cc1a975a27f0c"},
		{"Asia/Dushanbe", "848f466449ef22"},
		{"Asia/Famagusta", "2ba671b3c31cb6"},
		{"Asia/Gaza", "e46bcf0835cc53"},
		{"Asia/Hebron", "25db6167191615"},
		{"Asia/Ho_Chi_Minh", "0a710a81f69c49"},
		{"Asia/Hong_Kong", "806bf04c18f6a4"},
		{"Asia/Hovd", "ebfe57946e637f"},
		{"Asia/Irkutsk", "db273b1692c372"},
		{"Asia/Jakarta", "3b647dcf9395e0"},
		{"Asia/Jayapura", "ba4c08f61f9802"},
		{"Asia/Jerusalem", "767f41f547937e"},
		{"Asia/Kabul", "d87850b9aacffb"},
		{"Asia/Kamchatka", "81a0aa047b39d0"},
		{"Asia/Karachi", "770c7a71851ba4"},
		{"Asia/Kathmandu", "456dcaa8af0137"},
		{"Asia/Khandyga", "5e4017671a41db"},
		{"Asia/Kolkata", "7091596ffa3d6f"},
		{"Asia/Krasnoyarsk", "6893f45dd7fdfe"},
		{"Asia/Kuala_Lumpur", "645a6c2bba2c55"},
		{"Asia/Kuching", "d67f3550bbde37"},
		{"Asia/Kuwait", "17365c6a10dafc"},
		{"Asia/Macau", "d9151dea2b8c8a"},
		{"Asia/Magadan", "1a8ae99a2a2782"},
		{"Asia/Makassar", "107dc25fa434d7"},
		{"Asia/Manila", "c84dc54b6fda75"},
		{"Asia/Muscat", "f118fca9ad29e0"},
		{"Asia/Nicosia", "2f4fa03de437b2"},
		{"Asia/Novokuznetsk", "ec2ef7946ba75f"},
		{"Asia/Novosibirsk", "9021e880363f57"},
		{"Asia/Omsk", "41964ca7923243"},
		{"Asia/Oral", "6ec81cffbf80b9"},
		{"Asia/Phnom_Penh", "1f6756a1868e73"},
		{"Asia/Pontianak", "115332ec7eb26e"},
		{"Asia/Pyongyang", "8ccdf05936ed8e"},
		{"Asia/Qatar", "796751c31e4ef1"},
		{"Asia/Qostanay", "2b144db819acdc"},
		{"Asia/Qyzylorda", "f9e0ffbe492c37"},
		{"Asia/Riyadh", "c80ce74da80b59"},
		{"Asia/Sakhalin", "535c83421f1810"},
		{"Asia/Samarkand", "70e109f9b0cfa1"},
		{"Asia/Seoul", "a8a5fc7a91cc57"},
		{"Asia/Shanghai", "95033d34e64e9c"},
		{"Asia/Singapore", "4b773e48c663e2"},
		{"Asia/Srednekolymsk", "81137566bab0aa"},
		{"Asia/Taipei", "2feef03a36ac28"},
		{"Asia/Tashkent", "10c5bc983a82d3"},
		{"Asia/Tbilisi", "ce4cfb47645041"},
		{"Asia/Tehran", "23f12ca4e3cb7e"},
		{"Asia/Thimphu", "dd81f18038123a"},
		{"Asia/Tokyo", "d03f5792f1d28c"},
		{"Asia/Tomsk", "107d73d97f72dc"},
		{"Asia/Ulaanbaatar", "5a0e82625e24dd"},
		{"Asia/Urumqi", "85f5a295ad098f"},
		{"Asia/Ust-Nera", "7bb0a84ba4d874"},
		{"Asia/Vientiane", "c286c8838348ef"},
		{"Asia/Vladivostok", "126d3da5a620e6"},
		{"Asia/Yakutsk", "85f2e910aedbe1"},
		{"Asia/Yangon", "e14f2876c53bbc"},
		{"Asia/Yekaterinburg", "2e7a548ea689b7"},
		{"Asia/Yerevan", "1e8a75a85f5ca4"},

		// Atlantic
		{"Atlantic/Azores", "c17e63e7a47795"},
		{"Atlantic/Bermuda", "f77d33c2e783f0"},
		{"Atlantic/Canary", "8d56d8ca5a18c0"},
		{"Atlantic/Cape_Verde", "c019f82891856a"},
		{"Atlantic/Faroe", "397861b050ce85"},
		{"Atlantic/Madeira", "686b7ddda917ec"},
		{"Atlantic/Reykjavik", "7528488f12ad40"},
		{"Atlantic/South_Georgia", "4d91d7e278f4f6"},
		{"Atlantic/St_Helena", "e85c65f96e215b"},
		{"Atlantic/Stanley", "4c05bf0a896f10"},

		// Australia
		{"Australia/Adelaide", "16a54973c8f7e4"},
		{"Australia/Brisbane", "3db6e5843d96eb"},
		{"Australia/Broken_Hill", "6ceb80996c2ea7"},
		{"Australia/Currie", "b1c1c1bbae20f2"},
		{"Australia/Darwin", "234962e6dc34e1"},
		{"Australia/Eucla", "5727028491e6cc"},
		{"Australia/Hobart", "81f6690a0c5889"},
		{"Australia/Lindeman", "f1d0d2f8294397"},
		{"Australia/Lord_Howe", "b26eee0e505b4e"},
		{"Australia/Melbourne", "a78ce7fd38bcc0"},
		{"Australia/Perth", "825144946a6bed"},
		{"Australia/Sydney", "b27b602aa501cf"},

		// Europe
		{"Europe/Amsterdam", "c147c70a2945af"},
		{"Europe/Andorra", "f7385fce07528f"},
		{"Europe/Astrakhan", "c501ce72071bd2"},
		{"Europe/Athens", "95cf9c255a8e74"},
		{"Europe/Belgrade", "d843670b72bb5c"},
		{"Europe/Berlin", "99f0b5e419f71f"},
		{"Europe/Bratislava", "5ab4a43c9d1bad"},
		{"Europe/Brussels", "09a81d6661e555"},
		{"Europe/Bucharest", "af7cd0f919c0be"},
		{"Europe/Budapest", "e96c0a8575b76a"},
		{"Europe/Busingen", "a1389738cacf24"},
		{"Europe/Chisinau", "84509ac07b3c9b"},
		{"Europe/Copenhagen", "9f73c87f1649db"},
		{"Europe/Dublin", "054efa5769e8f1"},
		{"Europe/Gibraltar", "23f9eb35910abb"},
		{"Europe/Guernsey", "2ec9443d6aaee9"},
		{"Europe/Helsinki", "916b96b26aae66"},
		{"Europe/Isle_of_Man", "4ee77649b42a15"},
		{"Europe/Istanbul", "29085b405b3618"},
		{"Europe/Jersey", "77cad2aa484fe7"},
		{"Europe/Kaliningrad", "f7caa0a98c2b38"},
		{"Europe/Kiev", "099d2adb7f9a18"},
		{"Europe/Kirov", "939cbccf4227ce"},
		{"Europe/Lisbon", "aecadecb62cd44"},
		{"Europe/Ljubljana", "c6ecaa19d2838d"},
		{"Europe/London", "d851b36dde7621"},
		{"Europe/Luxembourg", "00ce06071085d3"},
		{"Europe/Madrid", "078cabaede7e5f"},
		{"Europe/Malta", "3918de644fa33e"},
		{"Europe/Mariehamn", "80eda2e8002858"},
		{"Europe/Minsk", "6efc4d56df0ff4"},
		{"Europe/Monaco", "79da1c56a2a06e"},
		{"Europe/Moscow", "cf6de567b8374c"},
		{"Europe/Oslo", "7c5749b6432a7d"},
		{"Europe/Paris", "cc31b47c7e352b"},
		{"Europe/Podgorica", "f0ff7b88cdf831"},
		{"Europe/Prague", "abed61aed76ec4"},
		{"Europe/Riga", "e1a7cd365eb973"},
		{"Europe/Rome", "180c6afadacf78"},
		{"Europe/Samara", "74c81269112d24"},
		{"Europe/San_Marino", "b7519296d6d56b"},
		{"Europe/Sarajevo", "9f3e3d84410537"},
		{"Europe/Saratov", "f99ef02ad92d10"},
		{"Europe/Simferopol", "eecdfc63db1b02"},
		{"Europe/Skopje", "3b7d54d1d3d2b6"},
		{"Europe/Sofia", "5be14a299ae92c"},
		{"Europe/Stockholm", "e5a2653e2f236f"},
		{"Europe/Tallinn", "4ee3d567cef1bc"},
		{"Europe/Tirane", "0a3c659cc7bd71"},
		{"Europe/Ulyanovsk", "84e65db711cafd"},
		{"Europe/Uzhgorod", "e882de2f323936"},
		{"Europe/Vaduz", "8dfc0c7b949121"},
		{"Europe/Vatican", "b6a92c3787b544"},
		{"Europe/Vienna", "3db791fb847adf"},
		{"Europe/Vilnius", "90310add4f39b7"},
		{"Europe/Volgograd", "ef0ef114b72e8a"},
		{"Europe/Warsaw", "4b772ed0d26e19"},
		{"Europe/Zagreb", "22029f966ab85d"},
		{"Europe/Zaporozhye", "09c44406c582fd"},
		{"Europe/Zurich", "650575e62b1f4a"},

		// Indian
		{"Indian/Antananarivo", "2c45287ca13922"},
		{"Indian/Chagos", "46794869967e7c"},
		{"Indian/Christmas", "063a8f5b1f7fcf"},
		{"Indian/Cocos", "f6674d53ca79f9"},
		{"Indian/Comoro", "74f9d1524bfd29"},
		{"Indian/Kerguelen", "dad832611a00e1"},
		{"Indian/Mahe", "419803fd1d1dd2"},
		{"Indian/Maldives", "7a51d39fbbc432"},
		{"Indian/Mauritius", "9843bc0435d26b"},
		{"Indian/Mayotte", "b90734cfef0168"},
		{"Indian/Reunion", "5bdc6ddba56253"},

		// Pacific
		{"Pacific/Apia", "b122ecf95cdb6e"},
		{"Pacific/Auckland", "d73daf5e9fd958"},
		{"Pacific/Bougainville", "63b89424d179f6"},
		{"Pacific/Chatham", "b6a0c2c5d47edc"},
		{"Pacific/Chuuk", "d4ddfdfb7802b4"},
		{"Pacific/Easter", "521a5bf45da863"},
		{"Pacific/Efate", "57a73edc802908"},
		{"Pacific/Enderbury", "68b7bd614d0126"},
		{"Pacific/Fakaofo", "139d48448805d6"},
		{"Pacific/Fiji", "c305cdf6f1e0e6"},
		{"Pacific/Funafuti", "01087200ea0be6"},
		{"Pacific/Galapagos", "159614fcc9665c"},
		{"Pacific/Gambier", "293890809aa95d"},
		{"Pacific/Guadalcanal", "3ca9875e5644b7"},
		{"Pacific/Guam", "0ab36cecf779eb"},
		{"Pacific/Honolulu", "4c7b9db133f68a"},
		{"Pacific/Kiritimati", "11e0463b051610"},
		{"Pacific/Kosrae", "ff28fb2d76ad15"},
		{"Pacific/Kwajalein", "d5f17024c40f0e"},
		{"Pacific/Majuro", "041fdd204d42a1"},
		{"Pacific/Marquesas", "439523787d9424"},
		{"Pacific/Midway", "5dd10bf23bf7c3"},
		{"Pacific/Nauru", "94c06fa864ae8a"},
		{"Pacific/Niue", "2a769964df035d"},
		{"Pacific/Norfolk", "7c0a07fd08f5a9"},
		{"Pacific/Noumea", "954538862c9cc5"},
		{"Pacific/Pago_Pago", "89642594f786db"},
		{"Pacific/Palau", "ba881f5b96c08d"},
		{"Pacific/Pitcairn", "b0c1adb711d175"},
		{"Pacific/Pohnpei", "491cb00b907324"},
		{"Pacific/Port_Moresby", "1f32297e91f219"},
		{"Pacific/Rarotonga", "9a1047399b5485"},
		{"Pacific/Saipan", "43b465f39cfa68"},
		{"Pacific/Tahiti", "f1476a03bd08f8"},
		{"Pacific/Tarawa", "4c5bd835eeb6d9"},
		{"Pacific/Tongatapu", "e9c095f8e1ca6a"},
		{"Pacific/Wake", "2d650fe51ff027"},
		{"Pacific/Wallis", "156d2c8f696dea"},

		// Other
		{"N/A", "e2f79e5b60330b"},
		{"UTC", "7e5f76c94a635c"},
		{"UTC ", "7e5f76c94a635c"},
		{" UTC", "7e5f76c94a635c"},
		{" UTC ", "7e5f76c94a635c"},
		{"UTC\n", "7e5f76c94a635c"},
	}

	for _, tt := range tests {
		got := Obfuscate(tt.timezone)
		assert.Equal(t, tt.hash, got)
	}
}