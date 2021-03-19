module Treely
  class Configuration
    attr_reader :style

    def initialize
      @style = {
        :dot         => %{\x2e},
        :space       => %{\x20},
        :new_line    => %{\x0a},
        :bar         => %{│   },
        :indent      => %{    },
        :branch      => %{├── },
        :last_branch => %{└── }
      }
    end

    def transform_style!(&block)
      @style.transform_values!(&block)
    end
  end

  def self.configuration
    @configuration ||= Configuration.new
  end

  def self.configuration=(config)
    @configuration = config
  end

  def self.configure
    yield configuration
  end
end
