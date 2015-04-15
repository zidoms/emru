App.Models.Task = Backbone.Model.extend({
	defaults: {
		title: '',
		done: false
	},

	validate: function(attrs) {
		if (!attrs.title)
			return 'Task title is required';
	}
});

App.Models.List = Backbone.Model.extend({
	defaults: {
		name: ''
	},

	validate: function(attrs) {
		if (!attrs.name)
			return 'List name is required';
	}
});
