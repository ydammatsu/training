# 例
# 座標 [0,0] にいる Player に対して、
# 前後左右に歩くコマンドを実行する。

#
# Command のクラス
#
class Command
  attr_reader :parameter

  def initialize(parameter = 0)
    # 歩数
    @parameter = parameter
  end
end

class Forward < Command
  def execute(position) = position[1] += parameter
end

class Back < Command
  def execute(position) = position[1] -= parameter
end

class Right < Command
  def execute(position) = position[0] += parameter
end

class Left < Command
  def execute(position) = position[0] -= parameter
end

#
# Reciever クラス
#
class Player
  def initialize
    @commands = []
  end

  def add_command(command)
    @commands << command
  end

  def position
    position = [0, 0]
    @commands.each { |cmmand| cmmand.execute(position) }
    position
  end
end

#
# 実行してみる
#

# Player のインスタンスを生成
player = Player.new

# Player のインスタンスに対して Command のインスタンスを追加していく。
player.add_command Left.new(1)    # 左に1歩
player.add_command Forward.new(5) # 前に5歩
player.add_command Right.new(3)   # 右に3歩

# [2, 5] と表示されるはず
p player.position
