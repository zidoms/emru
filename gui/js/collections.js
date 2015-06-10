App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task
});

App.Collections.Lists = Backbone.Collection.extend({
	model: App.Models.List,
	url: 'http://unix:/tmp/emru.sock:/lists',

	initialize: function() {
		this.fetch({
			success: this.fetchSuccess,
			error: this.fetchError
		});
	},

	fetchSuccess: function(collection, response) {
		console.log(collection.models);
	},

	fetchError: function(collection, response) {
		console.log('collection fetch error:');
		console.log(response);
	}
});
