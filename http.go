package main

//https://apis.scrimba.com/bored/documentation

// response
// {
//   "activity": "Learn Express.js",
//   "availability": 0.25,
//   "type": "education",
//   "participants": 1,
//   "price": 0.1,
//   "accessibility": "Few to no challenges",
//   "duration": "hours",
//   "kidFriendly": true,
//   "link": "https://expressjs.com/",
//   "key": "3943506"
// }

type boringResponse struct {
	Activity      string  `json:"activity"`
	Availability  float64 `json:"availability"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Accessibility string  `json:"accessibility"`
	Duration      string  `json:"duration"`
	KidFriendly   bool    `json:"kidfriendly"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
}
