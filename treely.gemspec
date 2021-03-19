Gem::Specification.new do |spec|
  spec.name       = 'treely'
  spec.version    = '1.0.1'
  spec.license    = 'MIT'
  spec.author     = 'minisode'
  spec.email      = 'minisode@189.cn'
  spec.executable = 'treely'
  spec.files      = Dir['lib/**/*']
  spec.summary    = 'A Ruby library for generating tree-like format'
  spec.homepage   = 'https://minisode.surge.sh/treely'

  spec.add_development_dependency 'minitest', '~> 5.11'
  spec.add_runtime_dependency 'activesupport', '~> 5.2'
  spec.add_runtime_dependency 'mercenary', '~> 0.3.6'
end
