# TCP DUMP MODULE IN RUBY WITH PACKETFU 
# CONTRIBUTION, CODE, ETC MENTIONS ARE LISTED AT THE START AND END OF THE 
# INITAL MAIN RED RABBIT VERSION 4 DOCUMENTATION 
#
#
#
require 'packetfu'
#
#############
class TCPdump
    attr_accessor :stats
    def initialize
      @stats = Hash.new(0)
      @stats[:connections] = []
    end
    def count_up
      @stats[:count] += 1
    end
    def count_connections
      @stats[:uniq] = @stats[:connections].count
    end
    # process the inital connections and incoming data
    def process_connection(args)
      processed = false
      @stats[:connections].each_with_index do |connection,index|
        if connection[:source] == args[:source] and connection[:destination] == args[:destination]
          @stats[:connections][index][:count] += 1
          processed = true
        end
      end
      # if not processed
      unless processed
        @stats[:connections] << { :source => args[:source], :destination => args[:destination], :count => 1 } 
      end
      count_up
      count_connections
      @stats[:connections]
    end
  end