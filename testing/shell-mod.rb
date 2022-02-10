
module Shell
    class Command
        attr_accessor :cmdline
        attr_reader   :output
        # initalizer
        def initialize(cmdline)
            @cmdline = cmdline.to_s
        end
        # execution 
        def execute
            @output = `#{@cmdline}`
            return @output
        end
    end
end

