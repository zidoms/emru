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
	}
});
