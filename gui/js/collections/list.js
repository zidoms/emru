// Merge list model and list collection
App.Models.List = Backbone.Model.extend({
	initialize: function() {
		this.tasks = new App.Collections.Tasks();
	},

	url: function() {
		url = 'http://unix:/tmp/emru.sock:/lists/';

		return url + this.get('name');
	},

	watch: function() {
		this.fetch({error: this.fErr});
	},

	parse: function(response) {
		if (!response || !response.tasks)
			return;

		tasks = response.tasks;
		for (var i = 0; i < tasks.length; i++) {
			task = new App.Models.Task(tasks[i]);
			task.set('list', this.get('name'));

			this.tasks.add(task);
		}
	},

	fErr: function(model, response) {
		console.log('model fetch err:');
		console.log(response);
	},
});

App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task,

	initialize: function() {
		this.on('add', this.added);
	},

	added: function(task) {
		task.view = new App.Views.Task({model: task}).render();
	},

	removeModel: function(i) {
		task = this.get(i);
		task.view.remove();
		this.remove(task);
	},

	clear: function() {
		while (task = this.first())
			task.destroy();
	}
});
