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
		var taskView = new App.Views.Task({ model: task });

		this.$el.append(taskView.render().el);
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

	done: function() {
		this.$el.addClass('done');
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

var tasksCollection = new App.Collections.Tasks();

var addTaskView = new App.Views.AddTask({ collection: tasksCollection });

var tasksView = new App.Views.Tasks({ collection: tasksCollection });
$('main').append(tasksView.render().el);

// TODO: Move to new view such as nav
$('.add').click(function(e) {
	$(this).parent('.action').toggleClass('active');
	$('#add').slideToggle(200);
	$('#add input').focus();
});

})();