package parsing

// AngleMap maps directions to angles
var AngleMap = map[interface{}]interface{}{
	'N': 90,
	'E': 180,
	'S': 270,
	'W': 0,
	90:  'N',
	180: 'E',
	270: 'S',
	0:   'W',
}
