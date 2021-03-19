require 'treely/configuration'
require 'treely/tree'

module Treely
  module_function

  def tree(source)
    Tree.new(source)
  end

  def [](name)
    configuration.send(name)
  end

  self.const_set 'version'.upcase, '1.0.1'
end
