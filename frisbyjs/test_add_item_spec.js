var frisby = require('frisby');
var rest = 'https://127.0.0.1:20443'

process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0"



for (i=500; i<520; i++) {

	frisby.create('post an item')
		.post(rest + "/item",
			{ name: "jasmine-"+i, description: "from jasmine" },
			{ json: true} )
		.expectStatus(200)
		.toss()
}

for (i=500; i<520; i++) {
	frisby.create('check items')
		.get(rest + "/items")
		.expectStatus(200)
		.expectJSON("jasmine-" +i,
			{name: "jasmine-" +i } )
		.toss()
}


