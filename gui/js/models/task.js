App.Models.Task = Backbone.Model.extend({
	defaults: {
		title: '',
		done: 'false'
	},

	url: function() {
		url = 'http://unix:/tmp/emru.sock:/lists/' + this.get('list') + '/tasks/';
		if (typeof this.get('id') !== 'undefined')
			url += this.get('id');

		return url;
	},

	toggle: function() {
		this.save({done: !this.get('done')});
	}
});
