//
// Todo Module
//

Todo = function() {

	// Add an item to a todo list
	// Assumes this is the todo list in question
	var add_item = function(data) {
		this.last().before(
			$("<li>").addClass("todo-item").attr("id","item-" + data.num).append(
				$("<span>").addClass("not-done")).append(
				$("<span>").addClass("desc").append(
					data.desc)));
	}

	// Toggle a todo list element to done
	// Assumes this is the li element in question
	var toggle_done = function() {
		this.find(".not-done").removeClass("not-done").addClass("done");
	}

	// Toggle a todo list element to not-done
	// Assumes this is the li element in question
	var toggle_not_done = function() {
		this.find(".done").removeClass("done").addClass("not-done");
	}

	var construct = function() {

	}

	return {
		init: construct
	};
}();
