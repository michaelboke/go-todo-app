//
// Todo Module
//

Todo = function() {

	// Add an item to a todo list
	// Assumes this is the todo list input in question
	var add_item = function(data) {
		this.parent().before(
			$("<li>").addClass("todo-item").attr("id","item-" + data.num).append(
				$("<span>").addClass("not-done")).append(
				$("<span>").addClass("desc").append(
					data.desc)));

		this.val('');
	}

	// Toggle a todo list element to done
	// Assumes this is the li element in question
	var toggle_done = function() {
		var li = this.parent();
		var done = !Boolean(li.data('done'));

		if done {
			li.removeClass("not-done").addClass("done");
		} else {
			li.removeClass("done").addClass("not-done");
		}

		$.ajax({
			url: "/done",
			method: "get",
			data: {
				num: li.data('num'),
				done: done
			}
		});
	}

	var construct = function() {
		$('#new-desc').change(add);
		$('.done-toggle').click(toggle_done);
	}

	return {
		init: construct
	};
}();
