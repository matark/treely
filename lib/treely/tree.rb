module Treely
  class Tree
    def initialize(source)
      @style = Treely[:style]
      @contents = %{}
      treely source
    end

    def to_s
      @contents
    end

    private

    def treely(source, *args)
      source.each.with_index do |item, index|
        @contents << args.map { |name| @style[name] }.join
        @contents << @style[index.zero? ? :dot : :bar]
        @contents << @style[:new_line]
        @contents << args.map { |name| @style[name] }.join
        @contents << @style[source[index.next] ? :branch : :last_branch]
        @contents << item.fetch(:name)
        @contents << @style[:new_line]

        if item.has_key?(:contents)
          treely item[:contents], *args, (source[index.next] ? :bar : :indent)
        end
      end
    end
  end
end
