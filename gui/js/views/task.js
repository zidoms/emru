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
		this.model.toggle();
	},

	removeTask: function() {
		this.model.destroy();
	},

	render: function() {
		this.$el.html(this.template(this.model.toJSON()));
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
