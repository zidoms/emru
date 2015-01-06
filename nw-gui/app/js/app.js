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
		done: false
	}
});

App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task,
	url: 'http://localhost:4040/tasks',
});

App.Views.Tasks = Backbone.View.extend({
	tagName: 'ul',

	id: 'list',

	initialize: function() {
		this.collection.on('add', this.addOne, this);

		$('main').append(this.render().el);
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

		this.model.set('done', !this.model.get('done')).save();
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

	initialize: function(attrs) {
		this.nav = attrs.nav;
	},

	events: {
		'submit': 'submit'
	},

	submit: function(e) {
		e.preventDefault();

		this.collection.create({
			title: $(e.currentTarget).find('input').val(),
			done: false
		}, {wait: true});

		this.nav.clearAdd();
	}
});

App.Views.Nav = Backbone.View.extend({
	el: '#actions',

	events: {
		'click .add': 'addTask'
	},

	addTask: function(e) {
		$(e.currentTarget).parent('.action').toggleClass('active');
		$('#add').slideToggle(200);
		$('#add input').focus();
	},

	clearAdd: function() {
		$('#add').slideToggle(200);
		$('.add').parent('.action').toggleClass('active');
		$('#add input').val('');
	}
});

var tasksCollection = new App.Collections.Tasks();
var appNav = new App.Views.Nav();

new App.Views.AddTask({collection: tasksCollection, nav: appNav});

window.App.List = new App.Views.Tasks({collection: tasksCollection});

})();