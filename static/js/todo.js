//
// Todo Module
//

Todo = function() {

	// Toggle a todo list element to done
	// @param e jQuery Event object
	var toggle_done = function(e) {
		var that = $(e.target);
		var li = that.parent();
		var done = !(li.data('done') === 'true');

		if (done) {
			li.removeClass("not-done").addClass("done");
			li.data('done', 'true');
		} else {
			li.removeClass("done").addClass("not-done");
			li.data('done', 'false');
		}

		$.ajax({
			url: "/done",
			method: "get",
			data: {
				num: li.data('num'),
				done: done
			}
		});
	};

	// Add an item to a todo list
	// @param e jQuery Event object
	var add_item = function(e) {
		var that = $(e.target);
		var desc = that.val();

		$.ajax({
			url: "/add",
			method: "get",
			data: {
				desc: desc
			},
			success: function(data) {
				that.parent().before(
					$("<li>").addClass("todo-item").addClass('not-done')
						.data('num',data.num).data('done', 'false').append(
					$("<input>").attr('type','checkbox')
						.addClass("done-toggle")).append(
					$("<p>").addClass("desc").append(
					desc)).on('click', toggle_done));
			},
			done: function() {
				that.val('');
			}
		});
	};

	var construct = function() {
		$('#new-desc').on('change', add_item);
		$('.done-toggle').on('click', toggle_done);
	};

	return {
		init: construct
	};
}();
