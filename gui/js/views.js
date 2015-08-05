App.Views.Task = Backbone.View.extend({
	tagName: 'li',
	className: 'task',
	template: template('task'),

	events: {
		'click .status': 'toggleStatus',
		'click .remove': 'removeTask'
	},

	initialize: function() {
		this.model.on('change', this.render, this);
		this.model.on('destroy', this.unrender, this);
	},

	toggleStatus: function(e) {
		this.model.set('done', !this.model.get('done')).save();
	},

	removeTask: function() {
		this.model.destroy();
	},

	render: function() {
		var template = this.template(this.model.toJSON());
		this.$el.html(template);
		if (isRTL(this.model.get('title')))
			this.$el.addClass('rtl');

		if (this.model.get('done') === true)
			this.$el.addClass('done');
		else
			this.$el.removeClass('done');

		$('#tasks').append(this.$el);

		return this;
	},

	unrender: function() {
		this.remove();
	}
});

App.Views.List = Backbone.View.extend({
	el: $('#list'),

	events: {
		'click .add': 'renderAdd',
		'click .clear': 'clear',
		'keyup input[name="task"]': 'keyup',
		'submit #add': 'addTask'
	},

	initialize: function() {
		this.before = null;

		this.model.on('change', 'render');
		this.render();
	},

	renderAdd: function() {
		$('#add input').val('');
		$('#add').toggle();
		$('#add input').focus();
	},

	keyup: function(event) {
		ch = $('#add input').val();
		if (isRTL(ch))
			$('#add input').attr('dir', 'rtl');
	},

	addTask: function(event) {
		url = 'http://unix:/tmp/emru.sock:/lists/' + this.model.get('name') + '/tasks';
		task = new App.Models.Task({title: $('#add input').val()});
		task.set('list', this.model.get('name'));
		task.save();

		this.renderAdd();
		event.preventDefault();
	},

	render: function() {
		if (this.before !== null) {
			tasks = this.before.tasks;
			for (var i = 0; i < tasks.models.length; i++)
				tasks.removeModel(i);
		}

		this.before = this.model;
		model = this.model;

		clearInterval(this.loop);
		this.loop = setInterval(function() { model.watch(); }, 500);
	},

	clear: function() {
		this.model.tasks.clear();
	}
});

App.Views.Lists = Backbone.View.extend({
	el: 'main',

	events: {
		'click #lists li a': 'changeList'
	},

	initialize: function() {
		this.Today = new App.Models.List({name: 'Today'});
		this.Week = new App.Models.List({name: 'Week'});
		this.Month = new App.Models.List({name: 'Month'});

		options = {type: 'post', url: 'http://unix:/tmp/emru.sock:/lists'};
		this.Today.save({}, options);
		this.Week.save({}, options);
		this.Month.save({}, options);

		this.listView = new App.Views.List({model: this.Today});
	},

	changeList: function(e) {
		el = $(e.target);
		prev = $('li.active');
		if (prev.text() == el.text()) return;

		prev.removeClass('active');
		el.parent('li').addClass('active');

		if (el.text() == 'Today')
			this.listView.model = this.Today;
		else if (el.text() == 'Week')
			this.listView.model = this.Week;
		else
			this.listView.model = this.Month;

		this.listView.render();
	}
});
