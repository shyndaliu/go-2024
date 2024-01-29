package api

type Titan struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Image     string   `json:"img"`
	Height    int      `json:"height"`
	Abilities []string `json:"abilities"`
	Ally      string   `json:"ally"`
}

var Titans = []Titan{
	{1, "Attack Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/a/ae/Attack_Titan_%28Anime%29_character_image_%28Eren_Jaeger%29.png/revision/latest/scale-to-width-down/350?cb=20170513212951", 15, []string{"Future memory inheritance"}, "Eldia"},
	{2, "Founding Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/3/3a/Eren_Jaeger_%28Anime%29_character_image_%28Founding_Titan%29.png/revision/latest/scale-to-width-down/350?cb=20220404091959", 350, []string{"Titan creation", "Titan behavioral control", "Memory and body manipulation of subjects of Ymir", "Telepathic communication with subjects of Ymir"}, "Eldia"},
	{3, "War Hammer Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/b/b2/Lara_Tybur_%28Anime%29_character_image_%28Titan%29.png/revision/latest/scale-to-width-down/350?cb=20210117211130", 15, []string{"Structural hardening", "Remote operation"}, "Eldia"},
	{4, "Colossal Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/7/78/Armin_Arlelt_%28Anime%29_character_image_%28Colossal_Titan%29.png/revision/latest/scale-to-width-down/350?cb=20220222211301", 60, []string{"Explosive transformation", "Immense size and strength", "Steam emission"}, "Eldia"},
	{5, "Armored Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/c/cf/Armored_Titan_%28Anime%29_character_image_%28Reiner_Braun%29.png/revision/latest/scale-to-width-down/350?cb=20170507070844", 15, []string{"Armored skin", "Hardening"}, "Marley"},
	{6, "Female Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/4/44/Female_Titan_%28Anime%29_character_image_%28Annie_Leonhart%29.png/revision/latest/scale-to-width-down/350?cb=20170708025046", 14, []string{"Versatility", "Titan", "Attraction", "Crystallization"}, "Marley"},
	{7, "Beast Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/8/89/Beast_Titan_%28Anime%29_character_image_%28Zeke_Jaeger%29.png/revision/latest/scale-to-width-down/350?cb=20180821055620", 17, []string{"Powerful and accurate throwing", "Hardening", "Speech", "Can turn subjects of Ymir into titans that he can crudely control (users with royal blood only)"}, "None"},
	{8, "Jaw Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/f/f3/Falco_Grice_%28Anime%29_character_image_%28Jaw_Titan%29.png/revision/latest/scale-to-width-down/350?cb=20220320210215", 5, []string{"Powerful jaw strength", "Hardened claws", "Great speed and agility"}, "Marley"},
	{9, "Cart Titan", "https://static.wikia.nocookie.net/shingekinokyojin/images/2/2e/Cart_Titan_%28Anime%29_character_image_%28Pieck_Finger%29.png/revision/latest/scale-to-width-down/350?cb=20210322020257", 4, []string{"Quadrupedal form", "High endurance", "Great speed", "Speech"}, "Marley"},
}
