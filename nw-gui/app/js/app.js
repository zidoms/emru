(function() {

window.App = {
	Models: {},
	Collections: {},
	Views: {}
};

window.template = function(id) {
	return _.template($('#' + id + 'Template').html());
};

App.Models.Task = Backbone.Model.extend({
	defaults: {
		title: '',
		body: '',
		done: false,
		date: ''
	}
});

App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task
});

App.Views.Tasks = Backbone.View.extend({
	tagName: 'ul',

	id: 'list',

	initialize: function() {
		this.collection.on('add', this.addOne, this);
	},

	render: function() {
		this.collection.each(this.addOne, this);

		return this;
	},

	addOne: function(task) {
		this.$el.append(new App.Views.Task({model: task}).render().el);
	}
});

App.Views.Task = Backbone.View.extend({
	tagName: 'li',

	className: 'task',

	template: template('task'),

	initialize: function() {
		this.model.on('remove', this.remove, this);
	},

	events: {
		'click .status': 'done',
		'click .remove': 'remove'
	},

	done: function(e) {
		this.$el.toggleClass('done');
		$(e.currentTarget).parent('.action').toggleClass('active');
	},

	remove: function() {
		this.$el.remove();
	},

	render: function() {
		var template = this.template(this.model.toJSON());

		this.$el.html(template);

		return this;
	}
});

App.Views.AddTask = Backbone.View.extend({
	el: '#add',

	events: {
		'submit': 'submit'
	},

	submit: function(e) {
		e.preventDefault();

		var title = $(e.currentTarget).find('input').val();

		this.collection.add(new App.Models.Task({title: title}));

		$('#add').slideToggle(200);
		$('.add').parent('.action').toggleClass('active');
		$(e.currentTarget).find('input').val('');
	}
});

App.Views.Nav = Backbone.View.extend({
	el: '#actions',

	events: {
		'click .add': 'addTask'
	},

	addTask: function(e) {
		console.log('hi');
		$(e.currentTarget).parent('.action').toggleClass('active');
		$('#add').slideToggle(200);
		$('#add input').focus();
	}
});

var tasksCollection = new App.Collections.Tasks();

var navView = new App.Views.Nav();
var addTaskView = new App.Views.AddTask({collection: tasksCollection});

var tasksView = new App.Views.Tasks({collection: tasksCollection});
$('main').append(tasksView.render().el);

})();