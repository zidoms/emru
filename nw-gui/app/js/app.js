(function() {

window.App = {
	Models: {},
	Collections: {},
	Views: {},
	gui: require('nw.gui')
};

window.template = function(id) {
	return _.template($('#' + id + 'Template').html());
};

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

App.Collections.Tasks = Backbone.Collection.extend({
	model: App.Models.Task,
	url: 'http://localhost:4040/tasks',
});

App.Views.Tasks = Backbone.View.extend({
	tagName: 'ul',

	id: 'list',

	initialize: function() {
		this.collection.on('add', this.addOne, this);

		this.collection.fetch({success:this.start()});
	},

	start: function() {
		$('main').append(this.render().el);
		this.fetcher = setInterval(this.fetch.bind(this), 7000);
	},

	render: function() {
		this.collection.each(this.addOne, this);

		return this;
	},

	addOne: function(task) {
		this.$el.append(new App.Views.Task({model: task}).render().el);
	},

	fetch: function() {
		this.collection.fetch();
	},

	close: function() {
		clearInterval(this.fetcher);
	}
});

App.Views.Task = Backbone.View.extend({
	tagName: 'li',

	className: 'task',

	template: template('task'),

	initialize: function() {
		this.model.on('change', this.render, this);
		this.model.on('destroy', this.unrender, this);
	},

	events: {
		'click .status': 'toggleStatus',
		'click .remove': 'deleteTask'
	},

	toggleStatus: function(e) {
		this.model.set('done', !this.model.get('done')).save();
	},

	deleteTask: function() {
		this.model.destroy();
	},

	render: function() {
		var template = this.template(this.model.toJSON());
		this.$el.html(template);

		if (this.model.get('done') === true)
			this.$el.addClass('done');
		else
			this.$el.removeClass('done');

		return this;
	},

	unrender: function() {
		this.remove();
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

var tasksCollection = new App.Collections.Tasks(),
	appNav = new App.Views.Nav();

new App.Views.AddTask({collection: tasksCollection, nav: appNav});

window.App.List = new App.Views.Tasks({collection: tasksCollection});

var win = App.gui.Window.get(),
	tray = new App.gui.Tray({icon: 'dist/emru.png'}),
	menu = new App.gui.Menu(),
	showing = true;

win.on('close', function() {
	win.hide();
	showing = false;
});

menu.append(
	new App.gui.MenuItem({
		label: 'Quit',
		click: function() {
			App.List.close();
			App.gui.App.quit();
		},
	})
);
tray.menu = menu;

tray.on('click', function() {
	if (showing)
		win.hide();
	else
		win.show();
	showing = !showing;
});


})();