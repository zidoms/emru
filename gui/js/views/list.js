// TODO:
// listen to collection if a tasks is added or etc
// cache elements like $('#add input') and etc
// listen to collection change filter all etc
// trim new task title
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
