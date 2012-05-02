class Array
  def iterate!(&block)
    self.each_with_index { |val, i| self[i] = yield val }
  end
end
