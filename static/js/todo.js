
Todo = {};

Todo.init = (function(bootstrap){

    var Item = Backbone.Model.extend({
        urlRoot: "/todo/list",
        defaults: {
            desc: '',
            done: false
        }
    });

    var List = Backbone.Collection.extend({
        url: '/todo/list',
        model: Item
    });

    var ItemView = Backbone.View.extend({
        tagName: 'li',
        events: {
            'click input.done-toggle': 'done'
        },
        initialize: function() {
            _.bindAll(this, 'render', 'flip', 'unrender', 'done', 'remove');

            this.model.bind('change', this.flip);
            this.model.bind('done', this.render);
            this.model.bind('remove', this.unrender);
        },
        flip: function() {
            if (this.model.get('done')) {
                $('input',this.el).attr('checked','checked');
                $(this.el).removeClass('not-done').addClass('done');
            } else {
                $('input',this.el).removeAttr('checked');
                $(this.el).removeClass('done').addClass('not-done');
            }
        },
        render: function() {
            var input = $('<input>').addClass('done-toggle').attr('type','checkbox');
            var p = $('<p>').append(this.model.get('desc'));
            if (this.model.get('done')) {
                input.attr('checked','checked');
                $(this.el).addClass('done');
            } else {
                $(this.el).addClass('not-done');
            }

            $(this.el).append(input).append(p);
            return this; // for chaining
        },
        unrender: function() {
            $(this.el).remove();
        },
        done: function() {
            var swapped = {
                done: !this.model.get('done')
            };
            this.model.save(swapped);
        },
        remove: function() {
            this.model.delete();
        }
    });

    var ListView = Backbone.View.extend({    
        el: $('#todo-list'), // attaches `this.el` to an existing element.
        events: {
            'click button#add': 'addItem',
            'change input#add-text': 'addItem'
        },

        initialize: function(){
            // Bind 'this' to functions
            _.bindAll(this, 'render', 'addItem', 'appendItem'); 

            this.collection = new List();
            this.collection.bind('add', this.appendItem);

            this.collection.reset(bootstrap);

            this.render();
        },

        render: function(){
            var self = this;

            $(this.el).append("<ul></ul>");
            _(this.collection.models).each(function(item) {
                self.appendItem(item);
            }, this);
        },

        addItem: function() {
            var text = $('#add-text', this.el).val();
            if (this.checkText(text)) {
                var item = new Item({
                    desc: text
                });
                item.save({});
                this.collection.add(item);
                $('#add-text', this.el).val('');
                $('#error', this.el).html('');
            } else {
                $('#error', this.el).html('Must have text!');
            }
        },

        checkText: function(text) {
            return (text.length > 0);
        },

        appendItem: function(item) {
            var itemView = new ItemView({
                model: item
            });
            $('ul', this.el).append(itemView.render().el);
        }
    });

    listView = new ListView();      
});