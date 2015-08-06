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
