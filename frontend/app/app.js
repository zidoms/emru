var $ = require('jquery');
$('body').on('click', '.icon-add', function(e) {
	e.stopPropagation();
	$('#add').slideToggle(200);
});
$('body').on('keyup', '#add input', function(e) {
	e.stopPropagation();
	if (e.keyCode == 13) {
		newTask = $('.task:first-child').clone();
		newTask.find('.title').html($(this).val());
		$('#list').prepend(newTask);
		$('#add').slideToggle(200);
		$(this).val('');
	}
});
$('body').on('click', '.icon-edit', function(e) {
	e.stopPropagation();
	$(this).parent('.action').toggleClass('active');
	$(this).parents('.actions').toggleClass('active');
	$(this).parents('.main').next('.etc').slideToggle(200);
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
$('body').on('click', '.icon-move-up', function(e) {
	e.stopPropagation();
	var task = $(this).parents('.task'),
		swap = task.prev('.task');
	if (task.is(':first-child')) return;
	Swap(task, swap);
});
$('body').on('click', '.icon-move-down', function(e) {
	e.stopPropagation();
	var task = $(this).parents('.task'),
		swap = task.next('.task');
	if (task.is(':last-child')) return;
	Swap(task, swap);
});

function Swap (task, swap) {
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
