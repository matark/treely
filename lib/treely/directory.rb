module Treely
  class Directory
    def initialize(path)
      @pathname = Pathname(path)
      @directories, @files = [], []
      @ignore_pattern = /\A\.(\.?|\w+)\z/

      unless @pathname.directory?
        @contents = []
      end
    end

    def contents
      @contents ||= digger.call(@pathname)
      [*@contents, { directories: @directories.size, files: @files.size }]
    end

    private

    def digger
      lambda do |pathname|
        pathname.children.reject { |n| @ignore_pattern.match?(n.basename.to_s) }.map do |pathname|
          if pathname.directory?
            @directories << pathname
            { name: pathname.basename.to_s, contents: digger.call(pathname) }
          else
            @files << pathname
            { name: pathname.basename.to_s }
          end
        end.sort_by do |item|
          [
            item[:contents].class.to_s,
            item[:name]
          ]
        end
      end
    end
  end
end
