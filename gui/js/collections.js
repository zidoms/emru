App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task,

	initialize: function() {
		this.on('add', this.added);
	},

	added: function(task) {
		new App.Views.Task({model: task}).render();
	}
});
