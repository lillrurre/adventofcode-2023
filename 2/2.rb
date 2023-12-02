class Game
  attr_accessor :red, :green, :blue

  def initialize
    @red = 0
    @green = 0
    @blue = 0
  end
end

input = File.read("input/2")
re = /(\d+) (\w+)/

part1 = 0
part2 = 0

input.strip.split("\n").each_with_index do |s, i|
  g = Game.new
    s.scan(re).each do |match|
    n = match[0].to_i
    color = match[1]
    case
    when color == "red" && g.red < n
      g.red = n
    when color == "green" && g.green < n
      g.green = n
    when color == "blue" && g.blue < n
      g.blue = n
    end
  end

  if g.red <= 12 && g.green <= 13 && g.blue <= 14
    part1 += i + 1
  end
  part2 += g.red * g.green * g.blue
end

puts "[1] Result: #{part1}"
puts "[2] Result: #{part2}"