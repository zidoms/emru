var gui = require('nw.gui'), JDate = require('jalali-date'),
	jdate = new JDate(), persianJs = require('persianjs'),
	sprintf = require("sprintf-js").sprintf, $ = require('jquery'),
	request = require('request');

var win = gui.Window.get(),
	tray = new gui.Tray({icon: 'icon/32/emru.png'}),
	menu = new gui.Menu(),
	showing = true;

menu.append(
	new gui.MenuItem({
		label: 'Quit',
		click: function() {
			gui.App.quit();
		},
	})
);
tray.menu = menu;

win.on('close', function() {
	win.hide();
	showing = false;
});
tray.on('click', function() {
	if (showing)
		win.hide();
	else
		win.show();
	showing = !showing;
});

var task = '<li class="task"><div class="main"><div class="title">%s</div><ul class="actions"><li class="action"><a class="hover-link icon-done"></a></li><li class="action"><a class="hover-link icon-cancel"></a></li></ul></div></li>';

$('#title').html(persianJs(jdate.format('dddd DD MMMM')).englishNumber().toString());
$('body').on('click', '.icon-add', function(e) {
	e.stopPropagation();
	$(this).parent('.action').toggleClass('active');
	$('#add').slideToggle(200);
	$('#add input').focus();
});
$('body').on('keyup', '#add input', function(e) {
	e.stopPropagation();
	if (e.keyCode == 13) {
		var options = {
			uri: 'http://localhost:4040/tasks',
			method: 'POST',
			json: {
				"title": $(this).val()
			}
		};

		request(options, function(error, response, body) {
			if (!error && response.statusCode == 200) {
				$('#add').slideToggle(200);
				$('.icon-add').parent('.action').toggleClass('active');
				$(this).val('');

				return;
			}
			console.log(error, response, body);
		});
	}
});
$('body').on('click', '.icon-done', function(e) {
	e.stopPropagation();
	$(this).parent('.action').toggleClass('active');
	$(this).parents('.task').toggleClass('done');
});
$('body').on('click', '.icon-cancel', function(e) {
	e.stopPropagation();
	$(this).parents('.task').slideUp(300, function() {
		$(this).remove();
	});
});

function Swap(task, swap) {
	var cur   = task.html(),
		cClss = task.attr('class'),
		swp   = swap.html();
		sClss = swap.attr('class');
	task.html(swp);
	task.attr('class', sClss);
	swap.html(cur);
	swap.attr('class', cClss);
}

/* Dev mode stuff */
var gulp = require('gulp');
gulp.task('reload', function() {
	if (location) location.reload();
});

gulp.watch(['**/*', '!styles/sass/*'], ['reload']);
