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

App.Views.List = Backbone.View.extend({
	tagName: 'section',
	id: 'list',

	events: {},

	initialize: function() {},

	render: function() {},

	unrender: function() {}
});

App.Views.Lists = Backbone.View.extend({
	tagName: 'ul',
	id: 'lists',

	events: {},

	initialize: function() {},

	render: function() {},

	unrender: function() {}
});
