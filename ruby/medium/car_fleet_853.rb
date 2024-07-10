def car_fleet(target, position, speed)
    n = position.length
    cars = position.zip(speed).sort_by { |pos, spd| -pos }
    times = cars.map { |pos, spd| (target - pos) / spd.to_f }
    
    fleets = 0
    max_time = 0
  
    times.each do |time|
      if time > max_time
        fleets += 1
        max_time = time
      end
    end
  
    fleets
end

# Test cases
target1 = 12
position1 = [10, 8, 0, 5, 3]
speed1 = [2, 4, 1, 1, 3]

target2 = 10
position2 = [3]
speed2 = [3]

target3 = 100
position3 = [0, 2, 4]
speed3 = [4, 2, 1]

puts car_fleet(target1, position1, speed1) # Output: 3
puts car_fleet(target2, position2, speed2) # Output: 1
puts car_fleet(target3, position3, speed3) # Output: 1
  