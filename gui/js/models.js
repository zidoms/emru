App.Models.Task = Backbone.Model.extend({
	defaults: {
		id: 0,
		title: '',
		body: '',
		done: false,
		created_at: ''
	}
});

App.Models.List = Backbone.Model.extend({
	url: 'http://unix:/tmp/emru.sock:/lists/',

	initialize: function() {
		this.url = this.url + this.get('name');
		this.tasks = new App.Collections.Tasks();
	},

	watch: function() {
		this.fetch({error: this.fErr});
	},

	parse: function(response) {
		tasks = response.tasks;
		for (var i = 0; i < tasks.length; i++) {
			this.tasks.add(new App.Models.Task(tasks[i]));
		}
	},

	fErr: function(model, response) {
		console.log('model fetch err:');
		console.log(response);
	},
});
